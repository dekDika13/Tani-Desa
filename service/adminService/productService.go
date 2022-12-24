package adminService

import "Tani-Desa/dto/adminDto"

// TODO Create Products
func (s *adminService) CreateProduct(payloads adminDto.ProductRequest) error {
	return s.adminRepsitory.CreateProduct(payloads)
}

func (s *adminService) GetAllProducts(adminId uint) ([]adminDto.ProductDTO, error) {
	// product := []adminDto.ProductDTO{}

	return s.adminRepsitory.GetAllProducts(adminId)

	// for _,v := range res{
	// 	product = append(product, adminDto.ProductDTO{
	// 		ProductID: ,
	// 	})
	// }

}
