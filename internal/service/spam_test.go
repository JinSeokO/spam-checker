package service

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

var spamService = Spam{}

func TestSpam_isSpam(t *testing.T) {
	//type args struct {
	//	contents         string
	//	spamLinkDomains  []string
	//	redirectionDepth int
	//}
	//assert.Equal(t, false, true, "It should not be equal")
	log.Println(strings.Fields("test result expect"))

	//testCases := []args{
	//	{
	//		contents:         fmt.Sprintf("spam spam https://goo.gl/nVLutc"),
	//		spamLinkDomains:  []string{"http://www.filekok.com/main"},
	//		redirectionDepth: 1,
	//	},
	//}
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
