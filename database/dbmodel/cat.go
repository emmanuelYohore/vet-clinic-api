package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Cat struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Age       int `gorm:"type:int"`
	Breed     string
	Weigth    int     `gorm:"type:int"`
	Visits    []Visit `gorm:"foreignKey:CatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CatRepository interface {
	Create(cat *Cat) (*Cat, error)
	FindAll() ([]*Cat, error)
	FindById(id uint) (*Cat, error)
	Update(cat *Cat) (*Cat, error)
	Delete(id uint, cat *Cat) error
	CatHistory(catID uint) ([]Visit, error)
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Delete(id uint, cat *Cat) error {
	return r.db.Delete(cat, id).Error
}

func (r *catRepository) FindById(id uint) (*Cat, error) {
	var cat Cat
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil

}

func (r *catRepository) Update(cat *Cat) (*Cat, error) {
	if err := r.db.Save(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *catRepository) Create(cat *Cat) (*Cat, error) {
	if err := r.db.Create(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *catRepository) FindAll() ([]*Cat, error) {
	var cats []*Cat
	if err := r.db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *catRepository) CatHistory(catID uint) ([]Visit, error) {
	var visits []Visit
	if err := r.db.
		Preload("Treatments").Where("cat_id = ?", catID).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}
