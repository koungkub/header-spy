package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", ":1323")
}

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println(string(reqBody))
	}))

	e.Any("/*", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	log.Fatal(e.Start(viper.GetString("PORT")))
}
