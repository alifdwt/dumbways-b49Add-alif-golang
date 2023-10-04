package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type PaslonRepository interface {
	FindPaslons() ([]models.Paslon, error)
	GetPaslon(ID int) (models.Paslon, error)
	// CreatePaslon(paslon models.Paslon) error
	CreatePaslon(paslon models.Paslon) (models.Paslon, error)
	UpdatePaslon(paslon models.Paslon, ID int) (models.Paslon, error)
	DeletePaslon(paslon models.Paslon, ID int) (models.Paslon, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryPaslon(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPaslons() ([]models.Paslon, error) {
	var paslons []models.Paslon
	// err := r.db.Raw("SELECT * FROM paslons").Scan(&paslons).Error
	err := r.db.Preload("Voter").Find(&paslons).Error

	return paslons, err
}

func (r *repository) GetPaslon(ID int) (models.Paslon, error) {
	var paslon models.Paslon
	// err := r.db.Raw("SELECT * FROM paslons WHERE id=?", ID).Scan(&paslon).Error
	err := r.db.First(&paslon, ID).Error

	return paslon, err
}

func (r *repository) CreatePaslon(paslon models.Paslon) (models.Paslon, error) {
    err := r.db.Exec("INSERT INTO paslons(name,vision,image,posted_at,updated_at) VALUES (?,?,?,?,?)", paslon.Name, paslon.Vision, paslon.Image, paslon.PostedAt, paslon.UpdatedAt).Error
    
	return paslon, err
}

func (r *repository) UpdatePaslon(paslon models.Paslon, ID int) (models.Paslon, error) {
	err := r.db.Raw("UPDATE paslons SET name=?, vision=?, image=?, updated_at=? WHERE id=?", paslon.Name, paslon.Vision, paslon.Image, paslon.UpdatedAt, ID).Scan(&paslon).Error

	return paslon, err
}

func (r *repository) DeletePaslon(paslon models.Paslon, ID int) (models.Paslon, error) {
	err := r.db.Raw("Delete FROM paslons WHERE id=?", ID).Scan(&paslon).Error

	return paslon, err
}