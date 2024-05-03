package models

import "time"

type Member struct {
	ID           int
	NamaDepan    string `gorm:"type: varchar(255)"`
	NamaBelakang string `gorm:"type: varchar(255)"`
	Birthday     time.Time
	JoinDate     time.Time
	UserID       int    `gorm:"foreignKey:UserID"`
	LevelID      int    `gorm:"foreignKey:LevelID"`
	Level        Level  `gorm:"foreignKey:LevelID"`
	GenderID     int    `gorm:"foreignKey:GenderID"`
	Gender       Gender `gorm:"foreignKey:GenderID"`
}

func (Member) TableName() string {
	return "members"
}
