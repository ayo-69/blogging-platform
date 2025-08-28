package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Bio      string `json:"bio"`
}
