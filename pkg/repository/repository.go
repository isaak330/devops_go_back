package repository

import "github.com/jmoiron/sqlx"

type CharactersList interface {
}

type Character interface{}

type Repository struct {
	CharactersList
	Character
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
