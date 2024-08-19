package database

import (
	"backend/handler/security"
	"backend/handler/security/models"
	"backend/tracer"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type apiRepo struct {
	dbGorm *gorm.DB
}

func NewApiRepo(dbGorm *gorm.DB) security.IApiRepo {
	return apiRepo{dbGorm: dbGorm}
}

func (a apiRepo) InsertApi(ctx context.Context, req models.Api) (res models.Api, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: InsertApi")
	defer span.End()
	err = a.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (a apiRepo) SelectApiByID(ctx context.Context, id string) (res models.Api, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectApiByID")
	defer span.End()
	err = a.dbGorm.WithContext(ctx).Where("id = ?", id).First(&res).Error
	return
}

func (a apiRepo) SelectApiByName(ctx context.Context, name string) (res models.Api, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectApiByName")
	defer span.End()
	err = a.dbGorm.Where("name = ?", name).First(&res).Error
	return
}

func (a apiRepo) SelectApiByEndpoint(ctx context.Context, endpoint, method string) (res models.Api, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectApiByEndpoint")
	defer span.End()
	err = a.dbGorm.Where("endpoint = ? and method = ?", endpoint, method).First(&res).Error
	return
}

func (a apiRepo) SelectApisByServiceID(ctx context.Context, serviceID int) (res []models.Api, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectApisByServiceID")
	defer span.End()
	err = a.dbGorm.Where("service_id = ?", serviceID).Find(&res).Error
	return
}

func (a apiRepo) UpdateApi(ctx context.Context, req models.Api) (res models.Api, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: UpdateApi")
	defer span.End()
	err = a.dbGorm.Model(&models.Api{}).Where("id = ?", req.ID).Updates(req).First(&res).Error
	return
}
