package shortLink

type ShortLinkRepository interface {
	Create(originalLink, shortenLink string) error
	Get(shortenLink string) (string, error)
	Exists(originalLink string) (string, error)
}
