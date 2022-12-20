package shorturl

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

type RedisRepo struct {
	addr        string
	redisClient *redis.Client
	ctx         context.Context
}

func NewRedisRepo(host string) *RedisRepo {
	result := RedisRepo{
		addr: host,
		ctx:  context.Background(),
	}
	return &result
}

func (r *RedisRepo) StoreUrl(shortUrl ShortUrl) error {
	err := r.client().Set(r.ctx, urlKey(shortUrl.ShortId), shortUrl.Url, 0*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) Find(shortId string) string {
	result, err := r.client().Get(r.ctx, urlKey(shortId)).Result()
	if err != nil {
		return ""
	}
	return result
}

func (r *RedisRepo) LogAccess(shortId, _ string) error {
	return r.client().Incr(r.ctx, urlAccessKey(shortId)).Err()
}

func (r *RedisRepo) ShortUrlAccessStats(shortId string) (*AccessStats, error) {
	count, err := r.client().Get(r.ctx, urlAccessKey(shortId)).Int()
	if err != nil {
		return &AccessStats{Count: 0}, nil
	}

	return &AccessStats{
		Count: count,
	}, nil
}

func (r *RedisRepo) client() *redis.Client {
	if r.redisClient == nil {
		r.redisClient = redis.NewClient(&redis.Options{
			Addr: r.addr,
		})
	}
	err := r.redisClient.Ping(r.ctx).Err()
	if err != nil {
		fmt.Println(err)
	}

	return r.redisClient
}

func (r *RedisRepo) set(shortUrl ShortUrl) {
	r.client().Set(r.ctx, shortUrl.ShortId, shortUrl.Url, 0*time.Second)
}

func urlKey(shortId string) string {
	return fmt.Sprintf("shorturl.%s.url", shortId)
}

func urlAccessKey(shortId string) string {
	return fmt.Sprintf("shorturl.%s.access", shortId)
}
