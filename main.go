package main

import (
	"Tani-Desa/database"
	"Tani-Desa/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	e := echo.New()

	router.New(e, database.DB)

	port := os.Getenv("PORT")

	_ = e.Start(port)
}
