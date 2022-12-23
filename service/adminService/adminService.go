package adminService

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/repository/adminRepository"
	"Tani-Desa/utils"
)

type AdminService interface {
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
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
