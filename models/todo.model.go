package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID         uint           `json:"id" gorm:"primary_key"`
	Name       string         `json:"name" gorm:"not null"`
	Note       *string        `json:"note" gorm:""`
	IsComplete bool           `json:"is_complete" gorm:"default:false;not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
