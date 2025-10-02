type CourseRepository struct{}
package repositories

import (
	"gorm.io/gorm"
	"smartcampus/internal/models"
)

type CourseRepository struct {
	DB *gorm.DB
}

func (r *CourseRepository) FindAll() ([]models.Course, error) {
	var courses []models.Course
	err := r.DB.Find(&courses).Error
	return courses, err
}

func (r *CourseRepository) FindByID(id string) (*models.Course, error) {
	var course models.Course
	err := r.DB.First(&course, "id = ?", id).Error
	return &course, err
}

func (r *CourseRepository) Create(course *models.Course) error {
	return r.DB.Create(course).Error
}

func (r *CourseRepository) Update(id string, updates map[string]interface{}) error {
	return r.DB.Model(&models.Course{}).Where("id = ?", id).Updates(updates).Error
}

func (r *CourseRepository) Delete(id string) error {
	return r.DB.Delete(&models.Course{}, "id = ?", id).Error
}

func (r *CourseRepository) FindByTeacherID(teacherID string, page, pageSize int) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64
	db := r.DB.Model(&models.Course{}).Where("teacher_id = ?", teacherID)
	db.Count(&total)
	err := db.Offset((page-1)*pageSize).Limit(pageSize).Find(&courses).Error
	return courses, total, err
}
