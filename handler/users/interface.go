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
	InsertUser(ctx context.Context, req models.User) (res models.User, err error)
	SelectUserByID(ctx context.Context, id int64) (res models.User, err error)
	SelectUserByEmail(ctx context.Context, email string) (res models.User, err error)
	SelectUserByPhone(ctx context.Context, phone string) (res models.User, err error)
	SelectUsersByID(ctx context.Context, id []int64) (res []models.User, err error)
	UpdateUser(ctx context.Context, req models.User) (res models.User, err error)
}

type IUserTypeRepo interface {
	SelectUserTypeByID(ctx context.Context, id int) (res models.UserType, err error)
}
