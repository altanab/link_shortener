package repository

import (
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"linkShortener/internal/pkg/shortLink"
)

type ShortLinkRep struct {
	DB *gorm.DB
}

func (rep *ShortLinkRep) Create(ol, sl string) error {
	link := shortLink.Link{
		OriginalLink: ol,
		ShortenLink: sl,
	}
	err := rep.DB.Table("links").Create(&link).Error
	if pgerr, ok := err.(*pgconn.PgError); ok {
		if pgerr.ConstraintName == "links_shorten_link_key" {
			return shortLink.LinkError{"shorten link already exists"}
		}
	}
	return err
}
func (rep *ShortLinkRep) Get(shortenLink string) (string, error) {
	var link shortLink.Link
	result := rep.DB.Table("links").
		Select("original_link").
		Where("shorten_link=?", shortenLink).
		Take(&link)
	if err := result.Error; err != nil {
		return "" , err
	} else if result.RowsAffected == 0 {
		return "", shortLink.LinkError{"link not found"}
	}
	return link.OriginalLink, nil
}
func (rep *ShortLinkRep) Exists(originalLink string) (string, error) {
	var link shortLink.Link
	result := rep.DB.Table("links").
		Select("shorten_link").
		Where("original_link=?", originalLink).
		Take(&link)
	if err := result.Error; err != nil {
		return "" , err
	} else if result.RowsAffected == 0 {
		return "", shortLink.LinkError{"link not found"}
	}
	return link.ShortenLink, nil
}