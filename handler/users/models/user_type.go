package models

import "time"

type UserType struct {
	Id        int64     `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserType) TableName() string {
	return "mst_user_type"
}

type UserTypeId int64

const (
	Root UserTypeId = 1
)

func (u UserTypeId) String() string {

	if u == Root {
		return "Root"
	}

	return ""

}
