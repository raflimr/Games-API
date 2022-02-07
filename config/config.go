package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	localUsername string = "root"
	localPassword string = ""
	localDatabase string = "db_games"
)

var (
	username string = os.Getenv("DB_USERNAME")
	password string = os.Getenv("DB_PASSWORD")
	database string = os.Getenv("DB_DATABASE")
	host     string = os.Getenv("DB_HOST")
	port     string = os.Getenv("DB_PORT")
)

var (
	dsn string
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	if host == "" {
		dsn = fmt.Sprintf("%v:%v@/%v?parseTime=true", localUsername, localPassword, localDatabase)
	} else {
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, port, database)
	}
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
