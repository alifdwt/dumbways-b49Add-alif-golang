package models

import (
	"time"
)

type Voter struct {
	ID int `json:"id" gorm:"primaryKey"`
	PaslonID int `json:"paslon_id"`
	VoterName string `json:"voter_name" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Paslon Paslon `json:"paslons"`
}

func (Voter) TableName() string {
	return "voters"
}