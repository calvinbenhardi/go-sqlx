package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres() *sqlx.DB {
	// TODO: remove hardcoded & use env
	db, err := sqlx.Connect("postgres", "port=5431 user=root password=secret dbname=bank sslmode=disable")
	if err != nil {
		fmt.Print(err)
	}

	return db
}
