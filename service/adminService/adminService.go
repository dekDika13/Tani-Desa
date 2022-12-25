package adminService

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/middleware"
	"Tani-Desa/repository/adminRepository"
	"Tani-Desa/utils"
	"errors"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error)
	//manage products
	CreateProduct(payloads adminDto.ProductRequest, file multipart.File) error
	GetAllProducts(adminId uint) ([]adminDto.ProductDTO, error)
	UpdateImageProduct(productId uint, file multipart.File) error
	GetProductById(adminId uint, productId uint) (adminDto.ProductDTO, error)
}

type adminService struct {
	adminRepsitory adminRepository.AdminRepository
}

func NewAdminService(adminRepo adminRepository.AdminRepository) *adminService {
	return &adminService{
		adminRepsitory: adminRepo,
	}
}

func (s *adminService) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	pw, err := utils.HashBcrypt(payloads.Password)

	if err != nil {
		return payloads, err
	}

	payloads.Password = pw

	res, err := s.adminRepsitory.RegisterAdmin(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

// TODO LOGIN ADMIN
func (s *adminService) LoginAdmin(payloads adminDto.LoginDTO) (adminDto.LoginJWT, error) {
	var temp adminDto.LoginJWT

	res, err := s.adminRepsitory.LoginAdmin(payloads)

	if errh := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(payloads.Password)); errh != nil {
		return temp, errors.New("username or password incorrect")
	}

	token, errt := middleware.CreateToken(res.AdminID, res.RoleId, res.Email)

	temp = adminDto.LoginJWT{
		Username: res.Username,
		Token:    token,
	}

	if err != nil {
		return temp, err
	}

	if errt != nil {
		return temp, errt
	}

	return temp, nil
}
