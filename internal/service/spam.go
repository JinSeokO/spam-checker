package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	RedirectionDepthLowerThanOneError = "redirectionDepth is should be over 0"
)

type Spam struct {
}

func (s Spam) IsSpam(contents string, spamLinkDomains []string, redirectionDepth int) (bool, error) {
	if redirectionDepth < 1 {
		return false, errors.New(RedirectionDepthLowerThanOneError)
	}

	isSpamRedirectionDomain := false
	lastPageRedirectionDepth := 0
	link, ok := s.hasURL(contents)

	if ok != true {
		return false, nil
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return false, err
	}

	client := new(http.Client)

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		lastPageRedirectionDepth = len(via)
		if lastPageRedirectionDepth >= redirectionDepth {
			redirectionDomain := via[redirectionDepth-1].URL.Host
			for _, spamLinkDomain := range spamLinkDomains {
				if spamLinkDomain == redirectionDomain {
					isSpamRedirectionDomain = true
					return nil
				}
			}
		}
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	lastPageRedirectionDepth += 1

	if isSpamRedirectionDomain {
		return true, nil
	}

	if lastPageRedirectionDepth == redirectionDepth {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		for _, spamLinkDomain := range spamLinkDomains {
			if strings.Contains(bodyString, fmt.Sprintf("<a href=\"%s\"></a>", spamLinkDomain)) {
				return true, nil
			}
		}
	}

	return false, nil
}

func (s Spam) hasURL(contents string) (url string, ok bool) {
	fields := strings.Fields(contents)

	for _, field := range fields {
		if s.isURL(field) {
			return field, true
		}
	}

	return "", false
}

func (s Spam) isURL(urlSample string) bool {
	_, err := url.ParseRequestURI(urlSample)
	if err != nil {
		return false
	}

	u, err := url.Parse(urlSample)

	if err != nil || u.Host == "" || u.Scheme == "" {
		return false
	}

	return true
}
