package models

import (
	"time"
)

type Api struct {
	ID        int64     `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Endpoint  string    `json:"endpoint" gorm:"column:endpoint"`
	Method    string    `json:"method" gorm:"column:method"`
	ServiceID int       `json:"service_id" gorm:"column:service_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Api) TableName() string {
	return "mst_api"
}
