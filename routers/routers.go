package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nba-simulation/handlers"
)

func Router(app *fiber.App) {

	app.Get("/", handlers.MatchStarting)

}
