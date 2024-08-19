package database

import (
	"backend/handler/users"
	"backend/handler/users/models"
	"backend/tracer"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	dbGorm *gorm.DB
	redis  *redis.Client
}

func NewUserRepo(dbGorm *gorm.DB, redis *redis.Client) users.IUserRepo {
	return userRepo{dbGorm: dbGorm, redis: redis}
}

func (u userRepo) InsertUser(ctx context.Context, req models.User) (res models.User, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: InsertUser")
	defer span.End()
	err = u.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (u userRepo) SelectUserByID(ctx context.Context, id int64) (res models.User, err error) {

	_, span := tracer.Tracer.Start(ctx, "Database: SelectUserByID")
	defer span.End()

	result, _ := u.redis.Get(context.Background(), fmt.Sprintf("user-id:%d", id)).Bytes()

	if result == nil {

		if err = u.dbGorm.Where("id = ?", id).First(&res).Error; err != nil {
			return
		}

		result, _ = json.Marshal(res)

		u.redis.Set(context.Background(), fmt.Sprintf("user-id:%d", id), result, 10*time.Second)

		return
	}

	if err = json.Unmarshal(result, &res); err != nil {
		return
	}

	return
}

func (u userRepo) SelectUserByEmail(ctx context.Context, email string) (res models.User, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectUserByEmail")
	defer span.End()
	err = u.dbGorm.Where("email = ?", email).First(&res).Error
	return
}

func (u userRepo) SelectUserByPhone(ctx context.Context, phone string) (res models.User, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectUserByPhone")
	defer span.End()
	err = u.dbGorm.Where("phone = ?", phone).First(&res).Error
	return
}

func (u userRepo) SelectUsersByID(ctx context.Context, id []int64) (res []models.User, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectUsersByID")
	defer span.End()
	err = u.dbGorm.Where("id in (?)", id).Find(&res).Error
	return
}

func (u userRepo) UpdateUser(ctx context.Context, req models.User) (res models.User, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: UpdateUser")
	defer span.End()
	err = u.dbGorm.Model(&models.User{}).Where("id = ?", req.ID).Updates(req).First(&res).Error
	return
}
