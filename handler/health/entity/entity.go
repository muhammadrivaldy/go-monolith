package database

import (
	"backend/config"
	health "backend/handler/health"
	db "backend/handler/health/entity/database"
)

func NewEntity(conf config.Configuration) (health.Entity, error) {

	database, err := db.NewDatabase(conf)
	if err != nil {
		return health.Entity{}, err
	}

	return health.Entity{
		Database: database}, nil
}
