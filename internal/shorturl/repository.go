package shorturl

import (
	"fmt"
	"shorturl/internal/params"
)

const TypeMemory string = "MEMORY"
const TypeMysql string = "MYSQL"

// Repository interface provides contract for repository to store short URL mapping
type Repository interface {
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

func NewRepo(parameters params.Params) Repository {
	repoType := parameters.GetWithDefault(params.Repository, TypeMemory)

	switch repoType {
	case TypeMemory:
		repo := NewShortUrlMemRepo()
		return &repo

	case TypeMysql:
		repo := NewMysqlRepo(
			parameters.Get("MYSQL_USER"),
			parameters.Get("MYSQL_PASS"),
			parameters.Get("MYSQL_HOST"),
			parameters.Get("MYSQL_DBNAME"),
			parameters.GetIntWithDefault("MYSQL_PORT", 3306))
		return &repo

	default:
		panic(fmt.Sprintf("unsupported repository '%s'", repoType))
	}
}
