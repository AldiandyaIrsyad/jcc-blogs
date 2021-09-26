package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	Privillege int64  `json:"privillege"`

	// Relation
	Posts    []Post
	Comments []Comment
}

// Privillege
// 1 = admin
// 2 = poster
// 3 = commenter
