package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"example.com/go-auth-globo/internal/config"
	_ "github.com/lib/pq"
)

var (
	host     string
	password string
	user     string
	port     string
	dbname   string
)

func OpenDB() (*sql.DB, error) {
	host = config.AppInfo.HOST
	password = config.AppInfo.PASSWORD
	user = config.AppInfo.USER
	port = config.AppInfo.PORT
	dbname = config.AppInfo.DBNAME

	var db *sql.DB
	var err error

	for i := range 2 {
		db, err = sql.Open("postgres", fmt.Sprintf("host=%s password=%s user=%s port=%s dbname=%s sslmode=disable",
			host,
			password,
			user,
			port,
			dbname,
		))

		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		slog.Info(fmt.Sprintf("Tentando conectar no banco... (%d/10)", i+1))
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		slog.Error("Error on connect database")
		panic(err)
	}

	return db, err
}
