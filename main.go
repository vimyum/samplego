package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	fmt.Printf("Start GO echo...")
	resBody := response{Message: "hello"}

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
		return c.JSON(http.StatusOK, resBody)
	})
	e.Logger.Fatal(e.Start(":8181"))
}
