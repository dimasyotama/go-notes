package main

import (
	"github.com/dimasyotama/go-notes/database"
	"github.com/dimasyotama/go-notes/middleware/ratelimiter"
	"github.com/dimasyotama/go-notes/redis"
	"github.com/dimasyotama/go-notes/router"
	"github.com/gofiber/fiber/v2"
	"github.com/dimasyotama/go-notes/middleware/logger"
)

func main(){
	// Start a new fiber app
	app := fiber.New()

	app.Use(ratelimiter.RateLimiter())
	app.Use(logger.CustomLogger)

	database.ConnectDB()
	redis.ConnectRedis()


	router.SetupRoutes(app)


	// Listen on PORT 1090
	app.Listen(":1090")
}