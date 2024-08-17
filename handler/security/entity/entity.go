package entity

import (
	"backend/handler/security"
	"backend/handler/security/entity/database"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type SecurityEntity struct {
	AccessRepo  security.IAccessRepo
	ApiRepo     security.IApiRepo
	ServiceRepo security.IServiceRepo
	VersionRepo security.IVersionRepo
}

func NewEntity(clientGorm *gorm.DB, clientRedis *redis.Client) (SecurityEntity, error) {
	return SecurityEntity{
		AccessRepo:  database.NewAccessRepo(clientGorm, clientRedis),
		ApiRepo:     database.NewApiRepo(clientGorm),
		ServiceRepo: database.NewServiceRepo(clientGorm),
		VersionRepo: database.NewVersionRepo(clientGorm),
	}, nil
}
