package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Title string `json:"title"`
}
