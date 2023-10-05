package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type PartyRepository interface {
	FindParties() ([]models.Party, error)
	GetParty(ID int) (models.Party, error)
	CreateParty(party models.Party) (models.Party, error)
}

func RepositoryParty(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindParties() ([]models.Party, error) {
	var parties []models.Party
	err := r.db.Preload("Paslon").Find(&parties).Error

	return parties, err
}

func (r *repository) GetParty(ID int) (models.Party, error) {
	var party models.Party
	err := r.db.Preload("Paslon").First(&party, ID).Error

	return party, err
}

func (r *repository) CreateParty(party models.Party) (models.Party, error) {
	err := r.db.Create(&party).Error

	return party, err
}