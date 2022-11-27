package repo

import (
	"errors"
	"shorturl/pkg/app"
)

// ShortUrlMemRepo ShortUrlRepo in-memory repository
//
// It stores short URL map in memory and as soon as application is stopped all maps are gone
// Use this repository only for testing purposes.
type ShortUrlMemRepo struct {
	urls map[string]string
}

func NewShortUrlMemRepo() ShortUrlMemRepo {
	return ShortUrlMemRepo{
		urls: make(map[string]string),
	}
}

func (u *ShortUrlMemRepo) StoreUrl(shortUrl app.ShortUrl) error {
	if u.urls[shortUrl.ShortId] != "" {
		return errors.New("already added")
	}
	u.urls[shortUrl.ShortId] = shortUrl.Url
	return nil
}

func (u *ShortUrlMemRepo) Find(shortId string) string {
	return u.urls[shortId]
}
