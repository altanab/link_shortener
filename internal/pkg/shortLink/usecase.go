package shortLink

type ShortLinkUsecase interface {
	Create(originalLink string) (string, error)
	Get(shortenLink string) (string, error)
}
