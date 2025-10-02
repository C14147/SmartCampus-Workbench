type AssignmentRepository struct{}
package repositories

import (
	"gorm.io/gorm"
	"smartcampus/internal/models"
)

type AssignmentRepository struct {
	DB *gorm.DB
}

func (r *AssignmentRepository) FindAll() ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.DB.Find(&assignments).Error
	return assignments, err
}

func (r *AssignmentRepository) FindByID(id string) (*models.Assignment, error) {
	var assignment models.Assignment
	err := r.DB.First(&assignment, "id = ?", id).Error
	return &assignment, err
}

func (r *AssignmentRepository) Create(assignment *models.Assignment) error {
	return r.DB.Create(assignment).Error
}

func (r *AssignmentRepository) Update(id string, updates map[string]interface{}) error {
	return r.DB.Model(&models.Assignment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *AssignmentRepository) Delete(id string) error {
	return r.DB.Delete(&models.Assignment{}, "id = ?", id).Error
}
