package database

import (
	"backend/handler/security"
	"backend/handler/security/models"
	"backend/tracer"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type serviceRepo struct {
	dbGorm *gorm.DB
}

func NewServiceRepo(dbGorm *gorm.DB) security.IServiceRepo {
	return serviceRepo{dbGorm: dbGorm}
}

func (s serviceRepo) InsertService(ctx context.Context, req models.Service) (res models.Service, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: InsertService")
	defer span.End()
	err = s.dbGorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&req).Error
	return req, err
}

func (s serviceRepo) SelectServiceByID(ctx context.Context, id int) (res models.Service, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectServiceByID")
	defer span.End()
	err = s.dbGorm.Where("id = ?", id).First(&res).Error
	return
}

func (s serviceRepo) SelectServiceByName(ctx context.Context, name string) (res models.Service, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectServiceByName")
	defer span.End()
	err = s.dbGorm.Where("name = ?", name).First(&res).Error
	return
}

func (s serviceRepo) SelectServices(ctx context.Context) (res []models.Service, err error) {
	_, span := tracer.Tracer.Start(ctx, "Database: SelectServices")
	defer span.End()
	err = s.dbGorm.Find(&res).Error
	return
}
