package settings

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var db *sql.DB

const (
	host     = "localhost"
	dbname   = "postgres"
	user     = "root"
	password = "root"
	schema   = "db_covid_palembang"
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", user, password, host, dbname, schema)
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(80)
	db.SetConnMaxLifetime(time.Nanosecond)
	db.SetMaxIdleConns(0)

	err = db.Ping()

	if err != nil {
		panic(err)
	}
}

func GetDatabase() *sql.DB {
	return db
}
