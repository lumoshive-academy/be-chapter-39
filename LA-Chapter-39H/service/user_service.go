package service

import (
	"golang-chapter-39/LA-Chapter-39H/models"
	"golang-chapter-39/LA-Chapter-39H/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(id uint) (models.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
