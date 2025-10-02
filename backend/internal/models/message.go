package models

import "time"

type Message struct {
    ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    SenderID    string    `gorm:"type:uuid;not null" json:"sender_id"`
    ReceiverID  string    `gorm:"type:uuid" json:"receiver_id"`
    CourseID    string    `gorm:"type:uuid" json:"course_id"`
    Title       string    `gorm:"size:200" json:"title"`
    Content     string    `gorm:"type:text;not null" json:"content"`
    MessageType string    `gorm:"size:50;default:notification" json:"message_type"`
    IsRead      bool      `gorm:"default:false" json:"is_read"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
