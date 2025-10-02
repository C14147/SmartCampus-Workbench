type MessageRepository struct{}
package repositories

import (
	"gorm.io/gorm"
	"smartcampus/internal/models"
)

type MessageRepository struct {
	DB *gorm.DB
}

func (r *MessageRepository) FindAll() ([]models.Message, error) {
	var messages []models.Message
	err := r.DB.Find(&messages).Error
	return messages, err
}

func (r *MessageRepository) FindByID(id string) (*models.Message, error) {
	var message models.Message
	err := r.DB.First(&message, "id = ?", id).Error
	return &message, err
}

func (r *MessageRepository) Create(message *models.Message) error {
	return r.DB.Create(message).Error
}

func (r *MessageRepository) Update(id string, updates map[string]interface{}) error {
	return r.DB.Model(&models.Message{}).Where("id = ?", id).Updates(updates).Error
}

func (r *MessageRepository) Delete(id string) error {
	return r.DB.Delete(&models.Message{}, "id = ?", id).Error
}
