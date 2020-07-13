package settings

import (
	"database/sql"
	"fmt"
)

type DatabaseConfig struct{}

const (
	host     = ""
	dbname   = ""
	user     = ""
	password = ""
	schema   = ""
)

// Postgresql Db Config
func (DatabaseConfig DatabaseConfig) GetDatabaseConfig() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", user, password, host, dbname, schema)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
