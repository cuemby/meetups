package models

import uuid "github.com/satori/go.uuid"

type Todo struct {
	Base
	Name        string    `gorm:"size:225;not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Priority    string    `gorm:"not null;size:225'" json:"priority"`
	UserId      uuid.UUID `gorm:"not null;" json:"user_id"`
}
