package service

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var spamService = Spam{}

func TestSpam_isSpam(t *testing.T) {
	type args struct {
		contents         string
		spamLinkDomains  []string
		redirectionDepth int
	}
	//assert.Equal(t, false, true, "It should not be equal")
	//spam, err := spamService.IsSpam("has url http://bit.ly/2yTkW52", []string{"bit.ly", "www.google.com"}, 1)
	spam, err := spamService.IsSpam("has url http://www.bit.ly/2yTkW52", []string{"bit.ly", "www.google.com"}, 2)
	if err != nil {
		t.Error(err)
	}
	log.Println(spam)
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "success http://www.bit.ly/2yTkW52",
			args: args{
				contents:         "has url http://www.bit.ly/2yTkW52",
				spamLinkDomains:  []string{"bit.ly", "www.google.com"},
				redirectionDepth: 2,
			},
			wantOk: true,
		},
		{
			name: "fail http://www.bit.ly/2yTkW52 cause depth",
			args: args{
				contents:         "has url http://www.bit.ly/2yTkW52",
				spamLinkDomains:  []string{"bit.ly", "www.google.com"},
				redirectionDepth: 1,
			},
			wantOk: false,
		},
		{
			name: "success http://bit.ly/2yTkW52",
			args: args{
				contents:         "has url http://bit.ly/2yTkW52",
				spamLinkDomains:  []string{"bit.ly", "www.google.com"},
				redirectionDepth: 1,
			},
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSpam, err := spamService.IsSpam(tt.args.contents, tt.args.spamLinkDomains, tt.args.redirectionDepth)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.wantOk, isSpam, "should be equal")
		})
	}
}

func Test_hasURL(t *testing.T) {
	type args struct {
		contents string
	}
	tests := []struct {
		name    string
		args    args
		wantUrl string
		wantOk  bool
	}{
		{
			name:    "default url",
			args:    args{contents: "has url https://www.google.com"},
			wantUrl: "https://www.google.com",
			wantOk:  true,
		},
		{
			name:    "has path param url",
			args:    args{contents: "has url https://www.google.com/test00"},
			wantUrl: "https://www.google.com/test00",
			wantOk:  true,
		},
		{
			name:    "no url case",
			args:    args{contents: "has url google"},
			wantUrl: "",
			wantOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUrl, gotOk := spamService.hasURL(tt.args.contents)
			assert.Equal(t, tt.wantOk, gotOk)
			assert.Equal(t, tt.wantUrl, gotUrl)
		})
	}
}
