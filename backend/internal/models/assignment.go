package models

import "time"

type Assignment struct {
    ID            string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    CourseID      string    `gorm:"type:uuid;not null" json:"course_id"`
    Title         string    `gorm:"size:200;not null" json:"title"`
    Description   string    `gorm:"type:text" json:"description"`
    DueDate       time.Time `json:"due_date"`
    MaxScore      int       `gorm:"default:100" json:"max_score"`
    AssignmentType string   `gorm:"size:50;default:homework" json:"assignment_type"`
    Attachments   string    `gorm:"type:jsonb" json:"attachments"`
    CreatedBy     string    `gorm:"type:uuid;not null" json:"created_by"`
    CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
