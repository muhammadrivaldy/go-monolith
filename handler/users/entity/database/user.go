package database

import (
	"backend/handler/users"
	"backend/handler/users/models"
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

func (u userRepo) InsertUser(req models.User) (res models.User, err error) {
	err = u.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (u userRepo) SelectUserById(id int64) (res models.User, err error) {

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

func (u userRepo) SelectUserByEmail(email string) (res models.User, err error) {
	err = u.dbGorm.Where("email = ?", email).First(&res).Error
	return
}

func (u userRepo) SelectUserByPhone(phone string) (res models.User, err error) {
	err = u.dbGorm.Where("phone = ?", phone).First(&res).Error
	return
}

func (u userRepo) SelectUsersById(id []int64) (res []models.User, err error) {
	err = u.dbGorm.Where("id in (?)", id).Find(&res).Error
	return
}

func (u userRepo) UpdateUser(req models.User) (res models.User, err error) {
	err = u.dbGorm.Model(&models.User{}).Where("id = ?", req.Id).Updates(req).First(&res).Error
	return
}
