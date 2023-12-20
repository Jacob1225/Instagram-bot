package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name string `json:"name" gorm:"text";not null;default:null`

}