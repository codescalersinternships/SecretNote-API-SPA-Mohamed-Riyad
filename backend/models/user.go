package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:TEXT;not null;unique" json:"user_name"`
	Password string `gorm:"type:TEXT;not null" json:"password"`
}
