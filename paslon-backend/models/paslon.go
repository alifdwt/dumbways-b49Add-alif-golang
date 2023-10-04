package models

import "time"

type Paslon struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	Vision string `json:"vision" gorm:"type:varchar(255)"`
	Image string `json:"image" gorm:"type:varchar(255)"`
	PostedAt time.Time `json:"posted_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Voter []Voter `json:"voters"`
}

func (Paslon) TableName() string {
    return "paslons" // Nama tabel yang digunakan di database
}