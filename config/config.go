package config

import (
	"github.com/emmanuelYohore/vet-clinic-api/database"
	"github.com/emmanuelYohore/vet-clinic-api/database/dbmodel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	CatRepository       dbmodel.CatRepository
	VisitRepository     dbmodel.VisitRepository
	TreatmentRepository dbmodel.TreatmentRepository
	UserRepository      dbmodel.UserRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	config.CatRepository = dbmodel.NewCatRepository(databaseSession)
	config.VisitRepository = dbmodel.NewVisitRepository(databaseSession)
	config.TreatmentRepository = dbmodel.NewTreatmentRipository(databaseSession)
	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	return &config, nil
}
