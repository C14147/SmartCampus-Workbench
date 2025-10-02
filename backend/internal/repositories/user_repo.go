type UserRepository struct{}
package repositories

import (
	"gorm.io/gorm"
	"smartcampus/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) Update(id string, updates map[string]interface{}) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.DB.Delete(&models.User{}, "id = ?", id).Error
}
