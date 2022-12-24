package adminController

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/middleware"
	"Tani-Desa/utils"
	"net/http"
	"strconv"

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

//TODO Get All Products

func (u *adminController) GetAllProducts(c echo.Context) error {
	adminID, _ := middleware.ClaimData(c, "adminID")
	conv_adminID := adminID.(float64)
	conv := uint(conv_adminID)

	res, err := u.adminServ.GetAllProducts(conv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Product Ditambahkan",
		Code:    http.StatusOK,
		Data:    res,
	})

}

// TODO Upload or Update Image
func (u *adminController) UpdateImageProduct(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}
	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()

	errs := u.adminServ.UpdateImageProduct(uint(convId), file)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: errs.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Successfully uploaded the file",
		Code:    http.StatusOK,
	})

}
