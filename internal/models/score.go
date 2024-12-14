package models

import "time"

type Score struct {
	ID       int `gorm:"primaryKey"`
	UserID   string
	GameID   int
	Score    int
	PlayedAt time.Time
	RoomID   int
}
