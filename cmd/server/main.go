package main

import (
	"ecommerce/internal/handler"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error { return handler.Index(c) })
	app.GET("/layout", func(c echo.Context) error { return handler.Layout(c) })

	err := app.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
