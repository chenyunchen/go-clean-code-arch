package provider

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/config"
)

func NewDBConn(config config.Config) (*sql.DB, error) {
	dbHost := config.Database.Host
	dbPort := config.Database.Port
	dbUser := config.Database.User
	dbPass := config.Database.Password
	dbName := config.Database.DBName
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "UTC")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
