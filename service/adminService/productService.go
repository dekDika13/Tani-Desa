package adminService

import "Tani-Desa/dto/adminDto"

// TODO Create Products
func (s *adminService) CreateProduct(payloads adminDto.ProductRequest) error {
	return s.adminRepsitory.CreateProduct(payloads)
}
