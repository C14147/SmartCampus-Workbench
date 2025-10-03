package models

import (
    "time"

    "gorm.io/gorm"
)

type Assignment struct {
    ID          string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    CourseID    string         `gorm:"type:uuid;not null" json:"course_id"`
    Title       string         `gorm:"size:200;not null" json:"title"`
    Description string         `json:"description"`
    AssignmentType string      `gorm:"size:50;default:'homework'" json:"assignment_type"`
    MaxScore    float64        `gorm:"type:decimal(5,2);default:100.00" json:"max_score"`
    DueDate     time.Time      `json:"due_date"`
    Attachments string         `gorm:"type:jsonb;default:'[]'" json:"attachments"`
    Status      string         `gorm:"size:20;default:'published'" json:"status"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
