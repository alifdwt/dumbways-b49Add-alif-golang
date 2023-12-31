package models

import (
	"time"
)

type Voter struct {
	ID int64 `json:"id" gorm:"primaryKey"`
	PaslonID int `json:"paslon_id"`
	UserID int `json:"user_id" gorm:"unique"`
	VoterName string `json:"voter_name" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Paslon Paslon `json:"paslons"`
}

func (Voter) TableName() string {
	return "voters"
}