package entity

import (
	"backend/handler/template"
	"backend/handler/template/entity/database"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type TemplateEntity struct {
	TemplateRepo template.ITemplateRepo
}

func NewTemplateEntity(clientGorm *gorm.DB, clientRedis *redis.Client) (TemplateEntity, error) {
	return TemplateEntity{
		TemplateRepo: database.NewTemplateRepo(clientGorm),
	}, nil
}
