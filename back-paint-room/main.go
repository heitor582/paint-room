package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/heitor582/paint_room/src/routes"
)

func main() {
	var PORT string = ":" + os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	routes.SetupTodoRoutes(app)
	app.Listen(PORT)
}
