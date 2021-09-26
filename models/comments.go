package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	Content string `json:"content" gorm:"type:text"`
}
