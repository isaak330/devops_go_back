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
	db, err := sqlx.Open("postgres", fmt.Printlf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfh.Port, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}
