package entity

import (
	"backend/handler/security"
	db "backend/handler/security/entity/database"
	"backend/models"
)

func NewEntity(conf models.Configuration) (security.Entity, error) {

	database, err := db.NewDatabase(conf)
	if err != nil {
		return security.Entity{}, err
	}

	return security.Entity{
		Database: database}, nil

}
