package router

import (
	"Tani-Desa/repository/adminRepository"
	"Tani-Desa/service/adminService"
	"Tani-Desa/utils"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	m.LogMiddleware(e)
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// TODO REPOSITORY
	adminRepository := adminRepository.NewAdminRepository(db)

	// TODO SERVICE
	adminService := adminService.NewAdminService(adminRepository)

	// TODO CONTROLLER
	adminController := adminController.NewAdminController(adminService)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "ping")
	})

	// TODO ADMIN ROUTE

	v1 := e.Group("/v1")
	v1.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
}
