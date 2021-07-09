package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	port := "8080"
	e := echo.New()
	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello", hello)
	e.POST("/users", createUser)
	e.GET("/users", findUser)
	e.PUT("/users", updateUser)
	e.DELETE("/users", deleteUser)

	e.Logger.Fatal(e.Start(":"+port))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func findUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
