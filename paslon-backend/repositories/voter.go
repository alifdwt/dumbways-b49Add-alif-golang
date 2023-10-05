package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type VoterRepository interface {
	FindVoters() ([]models.Voter, error)
	GetVoter(ID int64) (models.Voter, error)
	CreateVoter(voter models.Voter) (models.Voter, error)
	// UpdateVoter(voter models.Voter) (models.Voter, error)
	DeleteVoter(voter models.Voter) (models.Voter, error)
}

func RepositoryVoter(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindVoters() ([]models.Voter, error) {
	var voter []models.Voter
	err := r.db.Preload("Paslon").Find(&voter).Error
	
	return voter, err
}

func (r *repository) GetVoter(ID int64) (models.Voter, error) {
	var voter models.Voter
	err := r.db.Preload("Paslon").First(&voter, ID).Error

	return voter, err
}

func (r *repository) CreateVoter(voter models.Voter) (models.Voter, error) {
	err := r.db.Create(&voter).Error

	return voter, err
}

func (r *repository) DeleteVoter(voter models.Voter) (models.Voter, error) {
	err := r.db.Delete(&voter).Error

	return voter, err
}