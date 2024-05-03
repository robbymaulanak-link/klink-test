package models

import "time"

type Gender struct {
	ID        int    `gorm:"primaryKey"`
	Gender    string `gorm:"char(1)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Gender) TableName() string {
	return "genders"
}
