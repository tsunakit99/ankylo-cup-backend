package models

import "time"

type User struct {
    ID          string    `gorm:"primaryKey;size:128"`
    Email       string    `gorm:"size:255;uniqueIndex;not null"`
    DisplayName string    `gorm:"size:255;not null"`
    AvatarURL   string    `gorm:"type:text"`
    Coin        int32    // coin 0デフォルト
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
