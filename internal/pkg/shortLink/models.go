package shortLink

type Link struct {
	Id int
	OriginalLink string `gorm:"column:original_link"`
	ShortenLink string `gorm:"column:shorten_link"`
}

type LinkError struct {
	Message string
}

func (e LinkError) Error() string {
	return e.Message
}

