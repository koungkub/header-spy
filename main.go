package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	log.Fatal(e.Start(":1323"))
}
