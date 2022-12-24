package model

import (
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	BookingID uint           `gorm:"primaryKey;autoIncrement" json:"booking_id"`
	Products  Products       `gorm:"foreignKey:product_id"`
	Users     Users          `gorm:"foreignKey:user_id"`
	UserID    uint           `json:"user_id"`
	ProductID uint           `json:"product_id"`
	Status    string         `gorm:"size:13;not null" json:"status"`
	Start     string         `gorm:"size:20;not null" json:"start"`
	End       string         `gorm:"size:20;not null" json:"end"`
	CreatedAT time.Time      `json:"created_at"`
	UpdatedAT time.Time      `json:"updated_at"`
	DeletedAT gorm.DeletedAt `gorm:"index"`
}
