package main

import (
	"net/http"

	"github.com/gy-kim/plan-generator/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome Plan Generator")
	})
	e.POST("/generate-plan", handler.GeneratePlan)
	e.Logger.Fatal(e.Start(":8080"))
}
