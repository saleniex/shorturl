package repo

import "shorturl/pkg/app"

// ShortUrlRepo interface provides contract for repository to store short URL mapping
type ShortUrlRepo interface {
	// StoreUrl stores URL with given shortID
	//
	// Returns error != nil in case ShortId is already used or Url provided is invalid
	StoreUrl(shortUrl app.ShortUrl) error

	// Find finds Url by ShortId
	//
	// Returns Url being found. If URL is not found returns empty string
	Find(shortId string) string

	// LogAccess logs short URL access from remote IP address
	LogAccess(shortId, remoteIp string) error

	// ShortUrlAccessStats gathers access statistics for short URL
	//
	// Return AccessStats struct. If short ID is not found error is not nil
	ShortUrlAccessStats(shortId string) (*AccessStats, error)
}
