package router

import (
	m "Tani-Desa/middleware"
	"Tani-Desa/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	m.LogMiddleware(e)
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	
}
