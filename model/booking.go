package model

import (
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	BookingID uint           `gorm:"primaryKey;autoIncrement" json:"booking_id"`
	Product   Products       `gorm:"foreignKey:ProductId"`
	ProductId uint           `json:"product_id"`
	Users     Users          `gorm:"foreignKey:UserID"`
	UserID    uint           `json:"user_id"`
	Status    string         `gorm:"size:13;not null" json:"status"`
	Start     string         `gorm:"size:20;not null" json:"start"`
	End       string         `gorm:"size:20;not null" json:"end"`
	CreatedAT time.Time      `json:"created_at"`
	UpdatedAT time.Time      `json:"updated_at"`
	DeletedAT gorm.DeletedAt `gorm:"index"`
}
