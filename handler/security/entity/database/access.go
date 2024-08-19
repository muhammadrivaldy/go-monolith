package database

import (
	"backend/handler/security"
	"backend/handler/security/models"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type accessRepo struct {
	dbGorm *gorm.DB
	redis  *redis.Client
}

func NewAccessRepo(dbGorm *gorm.DB, redis *redis.Client) security.IAccessRepo {
	return accessRepo{dbGorm: dbGorm, redis: redis}
}

func (a accessRepo) InsertAccess(req models.Access) (res models.Access, err error) {
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (a accessRepo) InsertAccesses(req []models.Access) (res []models.Access, err error) {
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Find(&res).Error
	return
}

func (a accessRepo) SelectAccessByFilter(ctx context.Context, req util.FilterQuery) (res models.Access, err error) {

	query, _, arguments := req.BuildQuery()

	result, _ := a.redis.Get(context.Background(), fmt.Sprintf("select-access:%s", fmt.Sprintf("api-id:%d_user-type:%d", arguments...))).Bytes()

	if result == nil {

		if err = a.dbGorm.Where(query, arguments...).First(&res).Error; err != nil {
			return
		}

		result, _ = json.Marshal(res)

		a.redis.Set(context.Background(), fmt.Sprintf("select-access:%s", fmt.Sprintf("api-id:%d_user-type:%d", arguments...)), result, 1*time.Minute)

		return
	}

	if err = json.Unmarshal(result, &res); err != nil {
		return
	}

	return

}

func (a accessRepo) UpdateAccess(req models.Access) (res models.Access, err error) {
	err = a.dbGorm.Model(&models.Access{}).Where(`id = ?`, req.Id).Updates(req).First(&res).Error
	return
}

func (a accessRepo) SelectAccessByUserType(userTypeId int) (res []models.Access, err error) {
	err = a.dbGorm.Where("user_type_id = ?", userTypeId).Find(&res).Error
	return
}

func (a accessRepo) DeleteAccessesByUserTypeIdAndApiId(userTypeId int, apiId []int) (err error) {
	err = a.dbGorm.Where("user_type_id = ? and api_id in (?)", userTypeId, apiId).Delete(&models.Access{}).Error
	return
}
