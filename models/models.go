package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"  validate:"required, min=2, max=100" gorm:"text" default:"not null"`
	Email string `json:"email" validate:"email, required" gorm:"text" default:"not null"`
	Phone string `json:"phone" validate:"required" gorm:"text" default:"not null"`
}
