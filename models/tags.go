package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
}
