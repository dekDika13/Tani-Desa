package router

import (
	"Tani-Desa/controller/adminController"
	m "Tani-Desa/middleware"
	"Tani-Desa/repository/adminRepository"
	"Tani-Desa/service/adminService"
	"Tani-Desa/utils"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e.POST("/auth/login", adminController.LoginAdmin)

	// TODO REGISTER
	v1.POST("/register", adminController.RegisterAdmin, m.Authorization)

	//TODO MANAGE PRODUCT
	v1_Product := v1.Group("/product")
	{
		v1_Product.POST("/", adminController.CreateProduct)
		v1_Product.GET("/", adminController.GetAllProducts)
		v1_Product.PUT("/:id", adminController.UpdateImageProduct)
		v1_Product.GET("/:id", adminController.GetProductById)
		v1_Product.DELETE("/:id", adminController.DeleteProductById)
	}
}
