package models

import "time"

type Access struct {
	Id         int       `gorm:"column:id"`
	UserTypeId int       `gorm:"column:user_type_id"`
	ApiId      int       `gorm:"column:api_id"`
	CreatedBy  int64     `gorm:"column:created_by"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (a *Access) TableName() string {
	return "mst_access"
}
