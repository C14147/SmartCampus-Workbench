type SubmissionRepository struct{}
package repositories

import (
	"gorm.io/gorm"
	"smartcampus/internal/models"
)

type SubmissionRepository struct {
	DB *gorm.DB
}

func (r *SubmissionRepository) FindAll() ([]models.Submission, error) {
	var submissions []models.Submission
	err := r.DB.Find(&submissions).Error
	return submissions, err
}

func (r *SubmissionRepository) FindByID(id string) (*models.Submission, error) {
	var submission models.Submission
	err := r.DB.First(&submission, "id = ?", id).Error
	return &submission, err
}

func (r *SubmissionRepository) Create(submission *models.Submission) error {
	return r.DB.Create(submission).Error
}

func (r *SubmissionRepository) Update(id string, updates map[string]interface{}) error {
	return r.DB.Model(&models.Submission{}).Where("id = ?", id).Updates(updates).Error
}

func (r *SubmissionRepository) Delete(id string) error {
	return r.DB.Delete(&models.Submission{}, "id = ?", id).Error
}
