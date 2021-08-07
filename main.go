package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/user/create", userCreate)
	
	e.Start(":8080")
}

func userCreate(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
