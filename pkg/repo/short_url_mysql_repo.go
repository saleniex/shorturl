package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"shorturl/pkg/app"
)

type ShortUrlMysqlRepo struct {
	dsn string
}

func NewShortUrlMysqlRepo(user, pass, host, dbname string, port int) ShortUrlMysqlRepo {
	return ShortUrlMysqlRepo{
		dsn: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname),
	}
}

func (s ShortUrlMysqlRepo) StoreUrl(shortUrl app.ShortUrl) error {
	con, err := sql.Open("mysql", s.dsn)
	if err != nil {
		return err
	}
	defer func() {
		_ = con.Close()
	}()
	_, err = con.Query(
		"INSERT INTO shorturl (short_id, Url, created_at) VALUES (?,?,NOW())",
		shortUrl.ShortId,
		shortUrl.Url)
	if err != nil {
		return err
	}
	return nil
}

type shortUrl struct {
	ShortId string `json:"short_id"`
	Url     string `json:"Url"`
}

func (s ShortUrlMysqlRepo) Find(shortId string) string {
	con, err := sql.Open("mysql", s.dsn)
	if err != nil {
		return ""
	}
	defer func() {
		_ = con.Close()
	}()
	row := con.QueryRow("SELECT Url FROM shorturl WHERE short_id = ?", shortId)
	var shortUrl shortUrl
	err = row.Scan(&shortUrl.Url)
	if err != nil {
		return ""
	}

	return shortUrl.Url
}
