package models

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"type: varchar(255);unique"`
	Password  string `gorm:"type: varchar(255);"`
	IsAdmin   bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Member    Member
}

func (User) TableName() string {
	return "users"
}
