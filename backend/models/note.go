package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID    uint   `gorm:"not null;index" json:"user_id"`
	Title     string `gorm:"type:TEXT;not null" json:"title"`
	Content   string `gorm:"type:TEXT;not null" json:"content"`
	ViewCount int    `gorm:"default:0" json:"view_count"`
}
