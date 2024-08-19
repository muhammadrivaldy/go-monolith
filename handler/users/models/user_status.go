package models

import "time"

type UserStatus struct {
	ID        int64     `gorm:"column:id"`
	Key       string    `gorm:"column:key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (u *UserStatus) TableName() string {
	return "mst_user_status"
}

type UserStatusID int64

const (
	NonActive UserStatusID = -1
	Active    UserStatusID = 1
)

func (u UserStatusID) IsActive() bool {
	return u == Active
}

func (u UserStatusID) String() string {

	if u == NonActive {
		return "Non Active"
	} else if u == Active {
		return "Active"
	}

	return ""

}
