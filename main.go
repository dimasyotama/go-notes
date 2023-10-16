package main

import(
	"github.com/gofiber/fiber/v2"
	"github.com/dimasyotama/go-notes/database"
	"github.com/dimasyotama/go-notes/router"
)

func main(){
	// Start a new fiber app
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)


	// Listen on PORT 1090
	app.Listen(":1090")
}