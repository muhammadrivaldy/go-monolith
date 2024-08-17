package entity

import (
	"backend/handler/users"
	"backend/handler/users/entity/database"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

type UserEntity struct {
	UserRepo     users.IUserRepo
	UserTypeRepo users.IUserTypeRepo
}

func NewEntity(clientGorm *gorm.DB, clientRedis *redis.Client) (UserEntity, error) {
	return UserEntity{
		UserRepo:     database.NewUserRepo(clientGorm, clientRedis),
		UserTypeRepo: database.NewUserTypeRepo(clientGorm),
	}, nil
}
