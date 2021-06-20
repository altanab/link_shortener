package usecase

import (
	"linkShortener/internal/pkg/shortLink"
	"math/rand"
)

type ShortLinkUC struct {
	Rep shortLink.ShortLinkRepository
	Alphabet []byte
	ShortenLinkLen int
}

func (uc *ShortLinkUC) Create(originalLink string) (string, error) {
	shortenLink, err := uc.Rep.Exists(originalLink)
	if err == nil {
		return shortenLink, nil
	}
	for {
		shortenLink = uc.randomString(uc.ShortenLinkLen)
		err = uc.Rep.Create(originalLink, shortenLink)
		//перебираем укороченные строки, пока успешно не создадим запись в бд
		if err == nil {
			break
		} else if _, ok := err.(shortLink.LinkError); !ok {
			return "", err
		}
	}
	return shortenLink, nil
}
func (uc *ShortLinkUC) Get(shortenLink string) (string, error){
	return uc.Rep.Get(shortenLink)
}

func (uc *ShortLinkUC) randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = uc.Alphabet[rand.Int63() % int64(len(uc.Alphabet))]
	}
	return string(b)
}