package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type returnValueType struct {
	IpAddresses []string
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		ips := make([]string, 0)
		ips = append(ips, c.Request().RemoteAddr)
		forwardedForHeader := c.Request().Header.Get("x-forwarded-for")
		if forwardedForHeader != "" {
			forwardedFor := strings.Split(forwardedForHeader, ", ")
			ips = append(ips, forwardedFor...)
		}

		returnValue := new(returnValueType)
		returnValue.IpAddresses = ips
		return c.JSON(http.StatusOK, returnValue)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
