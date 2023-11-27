package users

import (
	"backend/handler/users/models"
	"backend/handler/users/payload"
	"backend/util"
	"context"
)

// IUserUseCase is a interface for layer business
type IUserUseCase interface {
	GetUserById(ctx context.Context, req payload.RequestGetUserById) (res payload.ResponseGetUserById, errs util.Error)
}

type IUserRepo interface {
	InsertUser(req models.User) (res models.User, err error)
	SelectUserById(id int64) (res models.User, err error)
	SelectUserByEmail(email string) (res models.User, err error)
	SelectUserByPhone(phone string) (res models.User, err error)
	SelectUsersById(id []int64) (res []models.User, err error)
	UpdateUser(req models.User) (res models.User, err error)
}

type IUserTypeRepo interface {
	SelectUserTypeById(id int) (res models.UserType, err error)
}
