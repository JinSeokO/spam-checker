package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var spamService = Spam{}

func TestSpam_isSpam(t *testing.T) {
	type args struct {
		contents         string
		spamLinkDomains  []string
		redirectionDepth int
	}
	assert.Equal(t, false, true, "It should not be equal")
	testCases := []args{
		{
			contents:         fmt.Sprintf("spam spam https://goo.gl/nVLutc"),
			spamLinkDomains:  []string{"http://www.filekok.com/main"},
			redirectionDepth: 1,
		},
	}
}
