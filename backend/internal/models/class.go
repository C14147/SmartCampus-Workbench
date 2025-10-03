package models

import (
    "time"

    "gorm.io/gorm"
)

type Class struct {
    ID          string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    SchoolID    string         `gorm:"type:uuid;not null" json:"school_id"`
    Name        string         `gorm:"size:100;not null" json:"name"`
    Grade       string         `gorm:"size:50" json:"grade"`
    Classroom   string         `gorm:"size:50" json:"classroom"`
    Capacity    int            `gorm:"default:40" json:"capacity"`
    HeadTeacher string         `gorm:"type:uuid" json:"head_teacher_id"`
    Status      string         `gorm:"size:20;default:'active'" json:"status"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
