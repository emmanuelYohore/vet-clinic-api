package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Date        time.Time
	Motif       string `gorm:"varchar(255)"`
	Veterinaire string

	CatID      uint
	Cat        Cat         `gorm:"foreignKey:CatID"`
	Treatments []Treatment `gorm:"constraint:OnDelete:CASCADE;"`
}

type VisitRepository interface {
	Create(visit *Visit) (*Visit, error)
	FindAll() ([]*Visit, error)
	FindByCatID(catID uint) ([]Visit, error)
	FindById(id uint) (*Visit, error)
	Update(visit *Visit) (*Visit, error)
	Delete(id uint, visit *Visit) error
	FilterByMotifOrVeterinaire(motif string, veterinaire string) ([]Visit, error)
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Delete(id uint, visit *Visit) error {
	return r.db.Delete(visit, id).Error
}

func (r *visitRepository) FindById(id uint) (*Visit, error) {
	var visit Visit
	if err := r.db.First(&visit, id).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (r *visitRepository) Update(visit *Visit) (*Visit, error) {
	if err := r.db.Save(visit).Error; err != nil {
		return nil, err
	}
	return visit, nil
}

func (r *visitRepository) Create(visit *Visit) (*Visit, error) {
	if err := r.db.Create(visit).Error; err != nil {
		return nil, err
	}
	return visit, nil
}

func (r *visitRepository) FindAll() ([]*Visit, error) {
	var visits []*Visit
	if err := r.db.Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *visitRepository) FindByCatID(catID uint) ([]Visit, error) {
	var visits []Visit
	if err := r.db.
		Where("cat_id = ?", catID).
		Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *visitRepository) FilterByMotifOrVeterinaire(motif string, veterinaire string) ([]Visit, error) {
	var visits []Visit
	query := r.db.Model(&Visit{})
	if motif != "" {
		query = query.Where("motif = ?", motif)
	}
	if veterinaire != "" {
		query = query.Where("veterinaire = ?", veterinaire)
	}
	if err := query.Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}
