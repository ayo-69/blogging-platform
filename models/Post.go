package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	User_id string   `json:"user_id" gorm:"not null"`
	Title   string   `json:"title" gorm:"not null"`
	Content string   `json:"content" gorm:"not null"`
	Tags    []string `json:"tags"`
}
