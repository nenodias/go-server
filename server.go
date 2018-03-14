package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Endereco struct {
	Rua string `json:"rua" xml:"rua" form:"rua" query:"rua"`
	Numero string `json:"numero" xml:"numero" form:"numero" query:"numero"`
	Cidade string `json:"cidade" xml:"cidade" form:"cidade" query:"cidade"`
}

type Usuario struct {
	Nome  string `json:"nome" xml:"nome" form:"nome" query:"nome"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	Endereco Endereco `json:"endereco" xml:"endereco" form:"endereco" query:"endereco"`
}

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
	e.POST("/users", func (c echo.Context) error {
		u := new(Usuario)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u) 
	})
	e.Logger.Fatal(e.Start(":3001"))
}
