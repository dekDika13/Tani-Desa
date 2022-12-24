package adminRepository

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AdminRepository interface {
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (model.Admins, error)
	//Manage Products
	CreateProduct(payloads adminDto.ProductRequest, link string) error
	GetAllProducts(adminId uint) ([]adminDto.ProductDTO, error)
	UpdateImageProduct(productId uint, link string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO REGISTER ADMIN
func (u *adminRepository) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	if err := u.db.Create(&model.Admins{
		RoleId:    payloads.RoleId,
		Username:  payloads.Username,
		Email:     payloads.Email,
		Password:  payloads.Password,
		CreatedAT: time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}

// TODO LOGIN ADMIN
func (u *adminRepository) LoginAdmin(payloads adminDto.LoginDTO) (model.Admins, error) {
	var admin model.Admins

	query := u.db.Where("email = ?", payloads.Email).First(&admin)
	if query.Error != nil {
		return admin, query.Error
	}

	if query.RowsAffected < 1 {
		return admin, errors.New("email is incorrect")
	}

	return admin, nil
}
