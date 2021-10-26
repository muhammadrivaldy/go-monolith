package database

import (
	health "backend/handler/health"
	db "backend/handler/health/entity/database"
	"backend/models"
)

func NewEntity(conf models.Configuration) (health.Entity, error) {

	database, err := db.NewDatabase(conf)
	if err != nil {
		return health.Entity{}, err
	}

	return health.Entity{
		Database: database}, nil

}
