package service

import "golang-chapter-39/LA-Chapter-39H-I/repository"

type Service struct {
	User UserService
}

func NewService(repo repository.Repository) Service {
	return Service{
		User: NewUserService(repo.User),
	}
}
