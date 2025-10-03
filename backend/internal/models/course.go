package models

import (
    "time"

    "gorm.io/gorm"
)

type Course struct {
    ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name      string         `gorm:"size:200;not null" json:"name"`
    Code      string         `gorm:"size:50;uniqueIndex;not null" json:"code"`
    Description string       `json:"description"`
    Credit    int            `gorm:"default:1" json:"credit"`
    TeacherID string         `gorm:"type:uuid" json:"teacher_id"`
    ClassID   string         `gorm:"type:uuid" json:"class_id"`
    Schedule  string         `gorm:"type:jsonb" json:"schedule"`
    Room      string         `gorm:"size:50" json:"room"`
    Status    string         `gorm:"size:20;default:'active'" json:"status"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
