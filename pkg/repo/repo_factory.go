package repo

import (
	"fmt"
	"shorturl/pkg/params"
)

const TypeMemory string = "MEMORY"
const TypeMysql string = "MYSQL"

func NewRepo(parameters params.Params) ShortUrlRepo {
	repoType := parameters.GetWithDefault(params.Repository, TypeMemory)

	switch repoType {
	case TypeMemory:
		repo := NewShortUrlMemRepo()
		return &repo

	case TypeMysql:
		repo := NewShortUrlMysqlRepo(
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
