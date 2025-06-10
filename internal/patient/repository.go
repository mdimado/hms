package patient

import (
	"hospital-management/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *Repository) GetAll() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Preload("CreatedByUser").Find(&patients).Error
	return patients, err
}

func (r *Repository) GetByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.Preload("CreatedByUser").First(&patient, id).Error
	return &patient, err
}

func (r *Repository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}