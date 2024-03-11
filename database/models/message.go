package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content        string `gorm:"size:255"`
	RecipientPhone string `gorm:"size:20;unique"`
	SentStatus     bool   `gorm:"default:false"`
}
