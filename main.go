package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/masutech16/iksm_server/model"
)


func main() {
	model.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo")
	}) //この関数を返すものをrouterに書いていく
	e.Start(":"+model.Data.Port)
}
