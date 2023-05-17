package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Instantiate
	e := echo.New()

	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello there!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
