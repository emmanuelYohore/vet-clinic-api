package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Treatment struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	VisitID   uint       `json:"visit_id"`
	Visit     Visit      `gorm:"foreignKey:VisitID" json:"visit"`
}

type TreatmentRepository interface {
	Create(treatment *Treatment) (*Treatment, error)
	FindAll() ([]*Treatment, error)
	FindById(id uint) (*Treatment, error)
	Update(treatment *Treatment) (*Treatment, error)
	Delete(id uint, treatment *Treatment) error
	FindByVisitID(visitID uint) ([]Treatment, error)
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRipository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) Delete(id uint, treatment *Treatment) error {

	return r.db.Delete(treatment, id).Error
}

func (r *treatmentRepository) FindById(id uint) (*Treatment, error) {
	var treatment Treatment
	if err := r.db.First(&treatment, id).Error; err != nil {
		return nil, err
	}
	return &treatment, nil
}

func (r *treatmentRepository) Update(treatment *Treatment) (*Treatment, error) {
	if err := r.db.Save(treatment).Error; err != nil {
		return nil, err
	}
	return treatment, nil
}

func (r *treatmentRepository) Create(treatment *Treatment) (*Treatment, error) {
	if err := r.db.Create(treatment).Error; err != nil {
		return nil, err
	}
	return treatment, nil
}

func (r *treatmentRepository) FindAll() ([]*Treatment, error) {
	var treatments []*Treatment
	if err := r.db.Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}

func (r *treatmentRepository) FindByVisitID(visitID uint) ([]Treatment, error) {
	var treatments []Treatment
	if err := r.db.Where("visit_id = ?", visitID).Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}
