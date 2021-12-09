package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/heitor582/paint_room/src/controllers"
)

// SetupRoutes initialzize routes of app
func SetupTodoRoutes(app *fiber.App) {
	route := app.Group("/room")
	route.Post("/", controllers.GetListOfCanOfPaint)
}
