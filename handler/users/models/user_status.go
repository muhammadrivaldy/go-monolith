package models

import "time"

type UserStatus struct {
	Id        int64     `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserStatus) TableName() string {
	return "mst_user_status"
}

type UserStatusId int64

const (
	NonActive UserStatusId = -1
	Active    UserStatusId = 1
)

func (u UserStatusId) IsActive() bool {
	return u == Active
}

func (u UserStatusId) String() string {

	if u == NonActive {
		return "Non Active"
	} else if u == Active {
		return "Active"
	}

	return ""

}
