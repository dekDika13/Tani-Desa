package model

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ProductID   uint           `gorm:"primaryKey:autoIncrement" json:"product_id"`
	Admins      Admins         `gorm:"foreignKey:admin_id"`
	AdminID     uint           `json:"admin_id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Type        string         `gorm:"size:50;not null" json:"type"`
	Qty         uint           `json:"qty"`
	Price       uint           `json:"price"`
	Description string         `json:"Description"`
	Address     string         `json:"address"`
	Owner       string         `gorm:"size:50;not null" json:"owner"`
	CreatedAT   time.Time      `json:"created_at"`
	UpdatedAT   time.Time      `json:"updated_at"`
	DeletedAT   gorm.DeletedAt `gorm:"index"`
}
