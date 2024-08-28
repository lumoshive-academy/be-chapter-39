package service

import (
	"golang-chapter-39/LA-Chapter-39H-I/models"
	"golang-chapter-39/LA-Chapter-39H-I/repository"
)

type UserService interface {
	CreateUser(user models.User) error
	GetUser(id uint) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUser(id uint) (models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
