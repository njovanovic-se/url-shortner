package models

type Shortener struct {
	ID          int64  `json:"id"`
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
	UserId      string `json:"user_id"`
}
