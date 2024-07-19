package handler

import (
	"ecommerce/internal/templ"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return render(c, templ.Index())
}

func Layout(c echo.Context) error {
	return render(c, templ.Layout("Layout"))
}
