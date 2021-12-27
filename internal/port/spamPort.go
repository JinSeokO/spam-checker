package port

type SpamPort interface {
	IsSpam(contents string, spamLinkDomains []string, redirectionDepth int) (bool, error)
}
