package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't read dotenv file")
	}
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		c.Redirect(http.StatusPermanentRedirect, "https://github.com/khafidprayoga")
		return nil
	})

	server.Any("/profile", func(c echo.Context) error {
		data := map[string]interface{}{
			"age":      os.Getenv("AGE"),
			"location": os.Getenv("LOCATION"),
			"email":    os.Getenv("EMAIL"),
			"jobdesk":  "Backend Developer",
		}

		profile := map[string]interface{}{
			"success": true,
			"code":    http.StatusOK,
			"data":    data,
			"stacks": []string{
				"Node.js", "Go", "AWS", "Docker",
			},
			"activity": []string{
				"Debugging", "Learning new things",
				"Explore about Cloud Native",
				"Drinking Coffee",
				"Playing FPS Game",
			},
		}

		return c.JSON(http.StatusOK, profile)
	})
	server.Start(":3000")
}
