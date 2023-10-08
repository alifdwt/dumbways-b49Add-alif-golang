package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(Email string) (models.User, error)
	Register(user models.User) (models.User, error)
	Login() ([]models.User, error)
	// Check(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserByEmail(Email string) (models.User, error) {
	var user models.User
	err := r.db.Preload("Voter").Where("email = ?", Email).First(&user).Error
	return user, err
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Voter").Find(&users).Error

	return users, err
}