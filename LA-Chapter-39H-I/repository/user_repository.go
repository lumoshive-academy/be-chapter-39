package repository

import (
	"golang-chapter-39/LA-Chapter-39H-I/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUser(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) UpdateUser(user models.User) error {
	return r.db.Save(&user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
