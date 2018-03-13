package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Usar o gin para rodar o projeto com auto reload
	e := echo.New()
	e.Static("/static", "assets")
	e.File("/", "public/index.html")
	e.GET("/hello", func(c echo.Context) error {
		nome := c.QueryParam("nome")
		if nome == ""{
			nome = "Mundo"
		}
		return c.String(http.StatusOK, "Olá " + nome + " seja bem-vindo")
	})
	e.Logger.Fatal(e.Start(":3001"))
}
