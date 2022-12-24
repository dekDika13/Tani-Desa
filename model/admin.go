package model

import (
	"time"
)

type Admins struct {
	Role      Roles     `gorm:"foreignKey:RoleId"`
	AdminID   uint      `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	RoleId    uint      `json:"role_id"`
	Email     string    `gorm:"size:50;not null" json:"email"`
	Username  string    `gorm:"size:50;not null" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAT time.Time `json:"created_at"`
	Image     string    `json:"image"`
	Address   string    `json:"address"`
}
