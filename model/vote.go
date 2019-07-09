package model

import "github.com/jinzhu/gorm"

//modelo de base de datos pibote que pone los comentarios del usuario
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentid" gorm:"not null"`
	UserID    uint `json:"userid" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
