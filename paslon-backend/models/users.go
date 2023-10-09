package models

import "time"

type User struct {
	ID int `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	PostedAt time.Time `json:"posted_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Voter []Voter `json:"voter"`
}