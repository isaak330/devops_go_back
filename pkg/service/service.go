package service

import "back/pkg/repository"

type CharactersList interface {
}

type Character interface{}

type Service struct {
	CharactersList
	Character
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
