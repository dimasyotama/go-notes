package main

import(
	"github.com/gofiber/fiber/v2"
	"github.com/dimasyotama/go-notes/database"
	"github.com/dimasyotama/go-notes/router"
	"github.com/dimasyotama/go-notes/redis"
)

func main(){
	// Start a new fiber app
	app := fiber.New()

	database.ConnectDB()
	redis.ConnectRedis()

	router.SetupRoutes(app)


	// Listen on PORT 1090
	app.Listen(":1090")
}