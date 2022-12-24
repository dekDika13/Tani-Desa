package model

import (
	"time"

	"gorm.io/gorm"
)

type Carts struct {
	CartID    uint           `gorm:"primaryKey;not null" json:"cart_id"`
	Product   Products       `gorm:"foreignKey:ProductId"`
	ProductId uint           `json:"product_id"`
	Status    string         `gorm:"size:13" json:"status"`
	Qty       uint           `json:"qty"`
	CreatedAT time.Time      `jsom:"created_at"`
	UpdatedAT time.Time      `json:"updated_at"`
	DeletedAT gorm.DeletedAt `gorm:"index"`
}
