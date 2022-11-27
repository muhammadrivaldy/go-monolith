package models

import "time"

type Access struct {
	ID         int       `gorm:"column:id"`
	UserTypeID int       `gorm:"column:user_type_id"`
	ApiID      int       `gorm:"column:api_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (Access) TableName() string {
	return "mst_access"
}
