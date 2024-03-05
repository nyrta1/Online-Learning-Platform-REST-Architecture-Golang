package repository

import (
	"gorm.io/gorm"
	"online-learning-platform/internal/models"
)

type RoleRepo interface {
	GetByID(id uint) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (rr *RoleRepository) GetByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := rr.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *RoleRepository) GetByName(name string) (*models.Role, error) {
	var role models.Role
	if err := rr.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
