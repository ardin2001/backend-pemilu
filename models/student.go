package models

import (
	"time"
)

type Student struct {
	ID        string `gorm:"primarykey"`
	NIM       int    `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}