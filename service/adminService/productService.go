package adminService

import (
	"Tani-Desa/dto/adminDto"
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// TODO Create Products
func (s *adminService) CreateProduct(payloads adminDto.ProductRequest) error {
	return s.adminRepsitory.CreateProduct(payloads)
}

func (s *adminService) GetAllProducts(adminId uint) ([]adminDto.ProductDTO, error) {

	return s.adminRepsitory.GetAllProducts(adminId)

}

func (s *adminService) UpdateImageProduct(productId uint, file multipart.File) error {
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	ctx := context.Background()

	result, errs := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
	if errs != nil {
		return errs
	}

	return s.adminRepsitory.UpdateImageProduct(productId, result.SecureURL)
}
