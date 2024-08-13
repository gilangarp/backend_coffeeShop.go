package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func PostgresConnection() (*sqlx.DB, error){
	
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, dbname)

	return sqlx.Connect("postgres", config)
}