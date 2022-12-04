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

func (s ShortUrlMysqlRepo) LogAccess(shortId, remoteIp string) error {
	con, err := sql.Open("mysql", s.dsn)
	if err != nil {
		return err
	}
	defer func() {
		_ = con.Close()
	}()
	_, err = con.Exec(
		`
			INSERT INTO access_log (created_at, shorturl_id, ip)
			SELECT
				NOW() AS created_at,
				id AS shorturl_id,
				? AS ip
			FROM shorturl
			WHERE short_id = ?`,
		remoteIp, shortId)
	if err != nil {
		return err
	}

	return nil
}

type queryResult struct {
	AccessCount int `json:"access_count"`
}

func (s ShortUrlMysqlRepo) ShortUrlAccessStats(shortId string) (*AccessStats, error) {
	con, err := sql.Open("mysql", s.dsn)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = con.Close()
	}()
	row := con.QueryRow(
		`
			SELECT COUNT(*) access_count
			FROM access_log l
				LEFT JOIN shorturl s ON l.shorturl_id = s.id
			WHERE s.short_id = ?;`,
		shortId)

	var queryResult queryResult
	scanErr := row.Scan(&queryResult.AccessCount)
	if scanErr != nil {
		return nil, scanErr
	}
	stats := AccessStats{
		Count: queryResult.AccessCount,
	}
	return &stats, nil
}
