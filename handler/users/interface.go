package users

import (
	"backend/handler/users/models"
	"backend/handler/users/payload"
	"backend/util"
	"context"
)

// IUserUseCase is a interface for layer business
type IUserUseCase interface {
	GetUserByID(ctx context.Context, req payload.RequestGetUserByID) (res payload.ResponseGetUserByID, errs util.Error)
}

type IUserRepo interface {
	InsertUser(req models.User) (res models.User, err error)
	SelectUserByID(id int64) (res models.User, err error)
	SelectUserByEmail(email string) (res models.User, err error)
	SelectUserByPhone(phone string) (res models.User, err error)
	SelectUsersByID(id []int64) (res []models.User, err error)
	UpdateUser(req models.User) (res models.User, err error)
}

type IUserTypeRepo interface {
	SelectUserTypeByID(id int) (res models.UserType, err error)
}
