// middleware/logger.go
package logger

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

var logFile *os.File

func init() {
	// Create or open the log file
	var err error
	logFile, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error opening log file: " + err.Error())
	}
}

func CustomLogger(c *fiber.Ctx) error {
	// Load the "Asia/Jakarta" time zone
	jakartaLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		// Handle error gracefully, use UTC as fallback
		fmt.Println("Error loading time zone, using UTC instead:", err)
	} else {
		// Set Fiber's time format
		c.Append("Time-Format", jakartaLocation.String())
	}

	// Format time in "Asia/Jakarta" time zone
	logTime := time.Now().In(jakartaLocation).Format("2006-01-02 15:04:05 MST")
	logMessage := fmt.Sprintf("[%s] - %s %s\n", logTime, c.Method(), c.Path())

	// Write log message to the file
	_, err = logFile.WriteString(logMessage)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}

	// Continue processing the request
	return c.Next()
}
