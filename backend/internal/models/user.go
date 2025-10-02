package models

import "time"

type User struct {
    ID           string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Username     string    `gorm:"unique;size:50;not null" json:"username"`
    Email        string    `gorm:"unique;size:100;not null" json:"email"`
    PasswordHash string    `gorm:"size:255;not null" json:"password_hash"`
    Role         string    `gorm:"size:20;not null" json:"role"`
    FullName     string    `gorm:"size:100;not null" json:"full_name"`
    AvatarURL    string    `gorm:"size:255" json:"avatar_url"`
    IsActive     bool      `gorm:"default:true" json:"is_active"`
    LastLoginAt  time.Time `json:"last_login_at"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
