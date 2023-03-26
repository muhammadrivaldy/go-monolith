package models

import "time"

type UserStatus struct {
	Id        int       `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserStatus) TableName() string {
	return "mst_user_status"
}

type UserStatusId int

const (
	NonActive UserStatusId = -1
	Active    UserStatusId = 1
)

func (u UserStatusId) IsActive() bool {
	return u == Active
}
