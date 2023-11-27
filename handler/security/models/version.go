package models

import "time"

type Version struct {
	Id        int       `json:"id" gorm:"column:id"`
	Version   string    `json:"version" gorm:"column:version"`
	Support   bool      `json:"support" gorm:"column:support"`
	CreatedBy int64     `json:"created_by" gorm:"column:created_by"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedBy int64     `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (a *Version) TableName() string {
	return "mst_version"
}
