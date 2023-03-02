package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfile() ([]models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfile() ([]models.Profile, error) {
	var profile []models.Profile
	err := r.db.Find(&profile).Error
	return profile, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error
	return profile, err
}
