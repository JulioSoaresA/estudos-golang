package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	conexao := "user=postgres dbname=web_golang_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
