package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	PostID    int64  `json:"post_id"`
	CreatedAt string `json:"created_at"`
}
