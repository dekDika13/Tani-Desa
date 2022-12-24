package adminRepository

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/model"
	"time"
)

func (u *adminRepository) CreateProduct(payloads adminDto.ProductRequest) error {

	if err := u.db.Create(&model.Products{
		AdminId:     payloads.AdminID,
		Name:        payloads.Name,
		Type:        payloads.Type,
		Qty:         payloads.Qty,
		Price:       payloads.Price,
		Description: payloads.Description,
		Address:     payloads.Address,
		Owner:       payloads.Owner,
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}
