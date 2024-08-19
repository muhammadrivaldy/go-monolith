package models

import "time"

type Customer struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Phone     string    `gorm:"column:phone"`
	Address   string    `gorm:"column:address"`
	StoreID   int       `gorm:"column:store_id"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedBy int       `gorm:"column:updated_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *Customer) TableName() string {
	return "mst_customer"
}
