package shorturl

import (
	"errors"
	"fmt"
)

// CachedRepo repository adds Redis caching layer to any repository which implements Repo interface
// Repository does two things:
// - Duplicates all mutations in cache
// - Queries are first performed in cache and only then in base repository
type CachedRepo struct {
	repo  Repo
	cache *RedisRepo
}

func NewCachedRepo(repo Repo, cache *RedisRepo) *CachedRepo {
	return &CachedRepo{
		repo:  repo,
		cache: cache,
	}
}

func (c *CachedRepo) StoreUrl(shortUrl ShortUrl) error {
	_ = c.cache.StoreUrl(shortUrl)
	return c.repo.StoreUrl(shortUrl)
}

func (c *CachedRepo) Find(shortId string) string {
	cachedResult := c.cache.Find(shortId)
	if cachedResult != "" {
		return cachedResult
	}
	return c.repo.Find(shortId)
}

func (c *CachedRepo) LogAccess(shortId, remoteIp string) error {
	err := c.repo.LogAccess(shortId, remoteIp)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot log access because %s", err.Error()))
	}
	return c.cache.LogAccess(shortId, remoteIp)
}

func (c *CachedRepo) ShortUrlAccessStats(shortId string) (*AccessStats, error) {
	result, err := c.cache.ShortUrlAccessStats(shortId)
	if err != nil && result.Count > 0 {
		return result, nil
	}
	return c.repo.ShortUrlAccessStats(shortId)
}
