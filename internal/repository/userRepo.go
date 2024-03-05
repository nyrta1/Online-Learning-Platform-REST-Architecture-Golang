package repository

import (
	"gorm.io/gorm"
	"online-learning-platform/internal/models"
)

type UserRepo interface {
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUser(id uint) error
	CreateUser(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) DeleteUser(id uint) error {
	var user models.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := ur.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
