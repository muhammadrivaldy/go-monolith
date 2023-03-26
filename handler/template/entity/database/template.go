package database

import (
	"backend/handler/template"

	"gorm.io/gorm"
)

type templateRepo struct {
	dbGorm *gorm.DB
}

func NewTemplateRepo(dbGorm *gorm.DB) template.ITemplateRepo {
	return templateRepo{dbGorm: dbGorm}
}
