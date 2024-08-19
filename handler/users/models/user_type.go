package models

import "time"

type UserType struct {
	ID        int64     `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserType) TableName() string {
	return "mst_user_type"
}

type UserTypeID int64

const (
	Root UserTypeID = 1
)

func (u UserTypeID) String() string {

	if u == Root {
		return "Root"
	}

	return ""

}
