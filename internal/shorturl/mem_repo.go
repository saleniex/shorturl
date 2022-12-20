package shorturl

import (
	"errors"
)

// MemoryRepo ShortUrlRepo in-memory repository
//
// It stores short URL map in memory and as soon as application is stopped all maps are gone
// Use this repository only for testing purposes.
type MemoryRepo struct {
	urls   map[string]string
	access map[string]map[string]int
}

func NewMemRepo() MemoryRepo {
	return MemoryRepo{
		urls: make(map[string]string),
	}
}

func (u *MemoryRepo) StoreUrl(shortUrl ShortUrl) error {
	if u.urls[shortUrl.ShortId] != "" {
		return errors.New("already added")
	}
	u.urls[shortUrl.ShortId] = shortUrl.Url
	return nil
}

func (u *MemoryRepo) Find(shortId string) string {
	return u.urls[shortId]
}

func (u *MemoryRepo) LogAccess(shortId, remoteIp string) error {
	if u.access[shortId] == nil {
		u.access[shortId] = make(map[string]int)
	}
	u.access[shortId][remoteIp] += 1
	return nil
}

func (u *MemoryRepo) ShortUrlAccessStats(shortId string) (*AccessStats, error) {
	if u.access[shortId] == nil {
		return nil, errors.New("No stats about " + shortId)
	}
	sum := 0
	for _, val := range u.access[shortId] {
		sum += val
	}
	stats := AccessStats{
		Count: sum,
	}
	return &stats, nil
}
