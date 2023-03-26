package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int64        `json:"id" gorm:"column:id"`
	Name       string       `json:"name" gorm:"column:name"`
	Phone      string       `json:"phone" gorm:"column:phone"`
	Email      string       `json:"email" gorm:"column:email"`
	Status     UserStatusId `json:"status" gorm:"column:status"`
	Password   string       `json:"-" gorm:"column:password"`
	UserTypeId UserTypeId   `json:"user_type_id" gorm:"column:user_type_id"`
	CreatedBy  int64        `json:"created_by" gorm:"column:created_by"`
	CreatedAt  time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedBy  int64        `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt  time.Time    `json:"updated_at" gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "mst_user"
}

func (u *User) EncryptPassword() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(hashPassword)
}

func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (u *User) IsExists() bool {
	return u.Id > 0
}
