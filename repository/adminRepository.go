package repository

import (
	admindto "Tani-Desa/dto/adminDto"
	"Tani-Desa/model"
	"errors"

	"gorm.io/gorm"
)

type AdminRepository interface {
	LoginAdmin(payloads admindto.LoginDTO) (model.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO LOGIN ADMIN
func (u *adminRepository) LoginAdmin(payloads admindto.LoginDTO) (model.Admin, error) {
	var admin model.Admin

	query := u.db.Where("username = ?", payloads.Username).First(&admin)
	if query.Error != nil {
		return admin, query.Error
	}

	if query.RowsAffected < 1 {
		return admin, errors.New("username is incorrect")
	}

	return admin, nil
}
