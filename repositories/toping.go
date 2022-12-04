package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type TopingRepository interface {
	FindTopings() ([]models.Toping, error)
	GetToping(ID int) (models.Toping, error)
	CreateToping(toping models.Toping) (models.Toping, error)
	DelToping(toping models.Toping) (models.Toping, error)
	UpdateToping(toping models.Toping) (models.Toping, error)
	UpdateStatus(status bool, ID int) error
}

func RepositoryToping(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindTopings() ([]models.Toping, error) {
	var topings []models.Toping
	err := r.db.Find(&topings).Error

	return topings, err
}
func (r *repository) GetToping(ID int) (models.Toping, error) {
	var toping models.Toping
	err := r.db.First(&toping, ID).Error

	return toping, err
}
func (r *repository) CreateToping(toping models.Toping) (models.Toping, error) {
	err := r.db.Create(&toping).Error

	return toping, err
}
func (r *repository) DelToping(toping models.Toping) (models.Toping, error) {
	err := r.db.Delete(&toping).Error

	return toping, err
}
func (r *repository) UpdateToping(toping models.Toping) (models.Toping, error) {
	err := r.db.Save(&toping).Error

	return toping, err
}
func (r *repository) UpdateStatus(status bool, ID int) error {
	var toping models.Toping
	r.db.First(&toping, ID)
	toping.Status = status
	err := r.db.Save(&toping).Error
	return err
}
