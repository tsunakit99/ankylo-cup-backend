package models

import "time"

type Room struct {
	ID        int `gorm:"primaryKey"`
	GameID    int
	Player1ID string
	Player2ID string
	Status    string `gorm:"size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
