package models

import "time"

type Level struct {
	ID        int    `gorm:"primaryKey"`
	Level     string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Level) TableName() string {
	return "levels"
}
