package models

import "time"

type Game struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"size:255;not null"`
	MaxPlayers int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
