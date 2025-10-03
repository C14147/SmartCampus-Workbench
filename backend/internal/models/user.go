package models

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    ID           string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Username     string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
    Email        string         `gorm:"uniqueIndex;size:100" json:"email"`
    PasswordHash string         `gorm:"size:255;not null" json:"-"`
    Role         string         `gorm:"size:20;not null;default:'student'" json:"role"`
    CreatedAt    time.Time      `json:"created_at"`
    UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
