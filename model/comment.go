package model

import "github.com/jinzhu/gorm"

//modelo de base de datos comentario
type Comment struct {
	gorm.Model
	UserID   uint      `json:"userid"`
	ParentID uint      `json:"parentid"`
	Votes    int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVote  int8      `json:"hasvote" gorm:"-"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"comment,omitempty"`
}
