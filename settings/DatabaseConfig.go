package settings

import (
	"Golang-Echo-MVC-Pattern/constant"
	"Golang-Echo-MVC-Pattern/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

var db *pgxpool.Pool

type Database struct{}

func init() {
	err := godotenv.Load()
	if err != nil {
		println(constant.MessageEnvironment)
	}

	host := os.Getenv(constant.DBHost)
	dbname := os.Getenv(constant.DBName)
	user := os.Getenv(constant.DBUser)
	password := os.Getenv(constant.DBPassword)
	schema := os.Getenv(constant.DBSchema)

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", user, password, host, dbname, schema)
	db, err = pgxpool.Connect(context.Background(), psqlInfo)

	if err != nil {
		utils.LogError(err, utils.DetailFunction())
		os.Exit(1)
	}
}

func (Database Database) GetDatabase() *pgxpool.Pool {
	return db
}
