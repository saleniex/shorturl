package repo

import "shorturl/pkg/app"

// ShortUrlRepo interface provides contract for repository to store short URL mapping
type ShortUrlRepo interface {
	// StoreUrl stores URL with given shortID
	// Returns error != nil in case ShortId is already used or Url provided is invalid
	StoreUrl(shortUrl app.ShortUrl) error

	// Find finds Url by ShortId
	// Returns Url being found. If URL is not found returns empty string
	Find(shortId string) string
}
