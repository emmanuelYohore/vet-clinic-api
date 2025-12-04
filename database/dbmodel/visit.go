package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Date        time.Time  `json:"date"`
	Motif       string     `json:"motif"`
	Veterinaire string     `json:"veterinaire"`

	CatID      uint        `json:"cat_id"`
	Cat        Cat         `gorm:"foreignKey:CatID" json:"cat"`
	Treatments []Treatment `gorm:"constraint:OnDelete:CASCADE;" json:"treatments,omitempty"`
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
