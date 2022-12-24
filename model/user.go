package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	UserID    uint           `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Roles     Roles          `gorm:"foreignKey:role_id"`
	RoleID    uint           `json:"role_id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Address   string         `json:"address"`
	Image     string         `json:"image"`
	CreatedAT time.Time      `json:"created_at"`
	UpdatedAT time.Time      `json:"updated_at"`
	DeletedAT gorm.DeletedAt `gorm:"index"`
}
