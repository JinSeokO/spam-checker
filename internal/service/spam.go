package service

import (
	"log"
	"net/url"
	"strings"
)

type Spam struct {
}

func (s Spam) IsSpam(contents string, spamLinkDomains []string, redirectionDepth int) (bool, error) {
	link, ok := hasURL(contents)
	if ok != true {
		return false, nil
	}

	log.Println(link)

	return false, nil
}

func hasURL(contents string) (url string, ok bool) {
	fields := strings.Fields(contents)

	for _, field := range fields {
		if isURL(field) {
			return field, true
		}
	}

	return "", false
}

func isURL(urlSample string) bool {
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
