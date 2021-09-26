package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	Title      string     `json:"title"`
	Summary    string     `json:"summary"`
	Content    string     `json:"content" gorm:"type:text"`
	Categories []Category `gorm:"many2many:post_categories"`
	Tags       []Tag      `gorm:"many2many:post_tags"`
	Comments   []Comment
}
