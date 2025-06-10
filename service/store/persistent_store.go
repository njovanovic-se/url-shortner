package store

import (
	"context"
	"fmt"

	"github.com/njovanovic-se/url-shortner/models"
)

type URLShortenerRepository interface {
	Save(ctx context.Context, shortener *models.Shortener) error
	Load(ctx context.Context, shortLink int64) (*models.Shortener, error)
}

type urlShortenerRepositoryImpl struct {
	db *DB
}

var (
	repo = &urlShortenerRepositoryImpl{}
)

func NewUrlShortenerRepositoryImpl(db *DB) *urlShortenerRepositoryImpl {
	repo.db = db
	return repo
}

func Save(ctx context.Context, shortener *models.Shortener) error {
	query := `INSERT INTO links (user_id, original_link, short_link)
			VALUES ($1, $2, $3)`

	err := repo.db.QueryRowContext(ctx, query,
		shortener.UserId,
		shortener.OriginalUrl,
		shortener.ShortUrl).Scan(&shortener.ID)

	if err != nil {
		return fmt.Errorf("failed to store new short link for userID: %s", shortener.UserId)
	}

	return nil
}

func Load(ctx context.Context, short_link string) (string, error) {
	query := `SELECT user_id, original_link, short_link FROM links
			WHERE short_link = $1`

	err := repo.db.QueryRowContext(ctx, query, short_link)

	if err != nil {
		return "", fmt.Errorf("short link not found for value: %s", short_link)
	}

	return short_link, nil
}
