package models

import "time"

type Party struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	PaslonID int `json:"paslon_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Paslon Paslon `json:"paslons"`
}

func (Party) TableName() string {
	return "parties"
}