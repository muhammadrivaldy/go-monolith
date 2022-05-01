package entity

import (
	"backend/config"
	"backend/handler/security"
	db "backend/handler/security/entity/database"
)

func NewEntity(conf config.Configuration) (security.Entity, error) {

	database, err := db.NewDatabase(conf)
	if err != nil {
		return security.Entity{}, err
	}

	return security.Entity{
		Database: database}, nil

}
