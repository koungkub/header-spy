package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("PORT", ":1323")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	log.Fatal(e.Start(viper.GetString("PORT")))
}
