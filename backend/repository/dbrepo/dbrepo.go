package dbrepo

import (
	"database/sql"
	"thefreepress/repository"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB) repository.DatabaseInterface {
	return &postgresDBRepo {
		DB: conn,
	}	
}