package middleware

import (
	"Tani-Desa/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(adminID uint, roleID uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["adminID"] = adminID
	claims["roleID"] = roleID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func CreateTokenUser(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	// claims["nik"] = nik
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func ClaimData(c echo.Context, field string) (interface{}, error) {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req := claims[field]

	return req, nil
}

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roleID, _ := ClaimData(c, "roleID")

		conv := fmt.Sprintf("%v", roleID)
		if conv != "1" {
			return c.JSON(http.StatusForbidden, utils.Response{
				Message: "you dont have permission to access",
				Code:    http.StatusForbidden,
			})
		}
		return next(c)
	}
}
