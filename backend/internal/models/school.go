package models

import (
    "time"

    "gorm.io/gorm"
)

type School struct {
    ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name      string         `gorm:"size:200;not null" json:"name"`
    Code      string         `gorm:"size:50;uniqueIndex;not null" json:"code"`
    Address   string         `json:"address"`
    Phone     string         `json:"phone"`
    Email     string         `json:"email"`
    Settings  string         `gorm:"type:jsonb;default:'{}'" json:"settings"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
