package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Interests []string `gorm:"type:text[]"` // Intereses del usuario
}

type Content struct {
	gorm.Model
	Title       string
	Description string
	Tags        []string `gorm:"type:text[]"` // Etiquetas para recomendar por similitud
}
