package service

type Spam struct {
}

func (s Spam) IsSpam(contents string, spamLinkDomains []string, redirectionDepth int) (bool, error) {
	//TODO implement me
	panic("implement me")
}
