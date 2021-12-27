package service

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Spam struct {
}

func (s Spam) IsSpam(contents string, spamLinkDomains []string, redirectionDepth int) (bool, error) {
	link, ok := s.hasURL(contents)
	if ok != true {
		return false, nil
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}

	client := new(http.Client)

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		log.Println(req)
		return nil
	}

	response, err := client.Do(req)
	if err != nil {
		return false, err
	}

	log.Println(response)

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
