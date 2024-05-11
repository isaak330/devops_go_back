package repository

import "github.com/jmoiron/sqlx"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Printlf())
}
