package models

import "time"

type UserType struct {
	Id        int       `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserType) TableName() string {
	return "mst_user_type"
}

type UserTypeId int

const (
	Root         UserTypeId = 1
	OwnerLaundry UserTypeId = 2
	AdminLaundry UserTypeId = 3
)
