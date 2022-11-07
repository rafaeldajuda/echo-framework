package main

import (
	"echo-framework/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// create a new echo instance
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(serverHeader)

	// routes
	e.GET("/cats/:data", handler.GetCats)
	e.POST("/cats", handler.AddCat)

	// start server
	e.Logger.Fatal(e.Start(":8000"))

}

// serverHeader - Custom Middleware
// middleware adds a custom header to the response.

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}
