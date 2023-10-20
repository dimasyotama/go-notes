package ratelimiter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

//func RateLimiter returns the rate limiting middleware
func RateLimiter() fiber.Handler{
	limiterConfig := limiter.Config{
		Max			:	10,
		Expiration	:	time.Minute,
		KeyGenerator: 	func(c *fiber.Ctx) string{
			return c.IP() //custom key generator based on IP address
		},

		LimitReached: func(c *fiber.Ctx) error{
			return c.Status(fiber.StatusTooManyRequests).JSON(
				fiber.Map{
					"status":"error",
					"message": "Rate Limit exceeded. Too many requests.",
				})
		},
	}
	return limiter.New(limiterConfig)
}