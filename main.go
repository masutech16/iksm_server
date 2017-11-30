package main

import (
	"net/http"
	"github.com/labstack/echo"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo")
	}) //この関数を返すものをrouterに書いていく
	e.Start(":1323")
}
