package adminController

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/middleware"
	"Tani-Desa/utils"
	"net/http"

	"github.com/labstack/echo"
)

// TODO Create Product
func (u *adminController) CreateProduct(c echo.Context) error {
	adminID, _ := middleware.ClaimData(c, "adminID")

	conv_adminID := adminID.(float64)
	conv := uint(conv_adminID)

	var payloads adminDto.ProductRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	temp := adminDto.ProductRequest{
		AdminID:     conv,
		Name:        payloads.Name,
		Type:        payloads.Type,
		Qty:         payloads.Qty,
		Price:       payloads.Price,
		Description: payloads.Description,
		Address:     payloads.Address,
		Owner:       payloads.Owner,
	}

	err := u.adminServ.CreateProduct(temp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Product Ditambahkan",
		Code:    http.StatusOK,
	})
}
