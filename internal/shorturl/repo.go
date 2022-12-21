package shorturl

import (
	"fmt"
	"shorturl/internal/params"
)

const TypeMemory string = "MEMORY"
const TypeMysql string = "MYSQL"
const TypeRedis string = "REDIS"
const TypeCachedMySql string = "CACHED_MYSQL"

// Repo interface provides contract for repository to store short URL mapping
type Repo interface {
	// StoreUrl stores URL with given shortID
	//
	// Returns error != nil in case ShortId is already used or Url provided is invalid
	StoreUrl(shortUrl ShortUrl) error

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

// NewRepo creates repository using provided parameters
func NewRepo(parameters params.Params) Repo {
	repoType := parameters.GetWithDefault(params.Repository, TypeMemory)

	switch repoType {
	case TypeMemory:
		return NewMemRepo()

	case TypeMysql:
		return newMySqlRepo(parameters)

	case TypeRedis:
		return newRedisRepo(parameters)

	case TypeCachedMySql:
		return NewCachedRepo(newMySqlRepo(parameters), newRedisRepo(parameters))

	default:
		panic(fmt.Sprintf("unsupported repository '%s'", repoType))
	}
}

func newMySqlRepo(params params.Params) *MysqlRepo {
	return NewMysqlRepo(
		params.Get("MYSQL_USER"),
		params.Get("MYSQL_PASS"),
		params.Get("MYSQL_HOST"),
		params.Get("MYSQL_DBNAME"),
		params.GetIntWithDefault("MYSQL_PORT", 3306))
}

func newRedisRepo(params params.Params) *RedisRepo {
	return NewRedisRepo(params.Get("REDIS_HOST"))
}
