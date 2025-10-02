package models

import "time"

type Course struct {
    ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name        string    `gorm:"size:200;not null" json:"name"`
    Description string    `gorm:"type:text" json:"description"`
    Code        string    `gorm:"unique;size:50;not null" json:"code"`
    TeacherID   string    `gorm:"type:uuid;not null" json:"teacher_id"`
    Semester    string    `gorm:"size:50" json:"semester"`
    Credits     int       `gorm:"default:0" json:"credits"`
    IsPublished bool      `gorm:"default:false" json:"is_published"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
