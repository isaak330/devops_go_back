package repository

type CharactersList interface {
}

type Character interface{}

type Repository struct {
	CharactersList
	Character
}

func NewRepository() *Repository {
	return &Repository{}
}
