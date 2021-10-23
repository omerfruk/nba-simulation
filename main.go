package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nba-simulation/database"
	"github.com/nba-simulation/routers"
)

func main() {
	fmt.Println("Backend Starting")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Content-Length, Accept-Encoding",
		MaxAge:           86400,
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE, UPDATE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	database.ConnectDB()

	//DB AutoMigrate and CreateRandomTeamAndMember
	//database.AutoMigrate()
	//database.CreateRandomTeamAndMember()

	routers.Router(app)
	port := "3000"

	fmt.Println("Backend Started")
	app.Listen(fmt.Sprintf(":%s", port))
}
