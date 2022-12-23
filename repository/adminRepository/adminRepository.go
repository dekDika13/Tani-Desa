package adminRepository

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/model"
	"time"

	"gorm.io/gorm"
)

type AdminRepository interface {
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO REGISTER ADMIN
func (u *adminRepository) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	if err := u.db.Create(&model.Admin{
		RoleId:    payloads.RoleId,
		Username:  payloads.Username,
		Password:  payloads.Password,
		CreatedAT: time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}
