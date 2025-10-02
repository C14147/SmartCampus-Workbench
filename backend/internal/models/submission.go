package models

import "time"

type Submission struct {
    ID           string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    AssignmentID string    `gorm:"type:uuid;not null" json:"assignment_id"`
    StudentID    string    `gorm:"type:uuid;not null" json:"student_id"`
    Content      string    `gorm:"type:text" json:"content"`
    Attachments  string    `gorm:"type:jsonb" json:"attachments"`
    SubmittedAt  time.Time `gorm:"autoCreateTime" json:"submitted_at"`
    Status       string    `gorm:"size:20;default:submitted" json:"status"`
    Score        int       `json:"score"`
    Feedback     string    `gorm:"type:text" json:"feedback"`
    GradedBy     string    `gorm:"type:uuid" json:"graded_by"`
    GradedAt     time.Time `json:"graded_at"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
