package models

import "time"

type Access struct {
	Id         int       `gorm:"column:id"`
	UserTypeId int       `gorm:"column:user_type_id"`
	ApiId      int       `gorm:"column:api_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (a *Access) TableName() string {
	return "mst_access"
}
