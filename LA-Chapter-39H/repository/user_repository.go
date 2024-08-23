package repository

import (
	"golang-chapter-39/LA-Chapter-39H/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUser(id uint) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUser(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) UpdateUser(user models.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
