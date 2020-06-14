package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os/exec"
	"github.com/google/uuid"
)

type response struct {
	Message string `json:"message"`
}

func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {
		username := c.Param("username")
		u, _ := uuid.NewRandom()
        uu := u.String()
		tmpFileNmae := uu + ".svg"
		cmdstr := fmt.Sprintf("curl https://github.com/%s | " + 
			"awk '/<svg.+class=\"js-calendar-graph-svg\"/,/svg>/' | " +
			"sed -e 's/<svg/<svg xmlns=\"http:\\/\\/www.w3.org\\/2000\\/svg\"/' > %s", username, tmpFileNmae);
		cmd := exec.Command("sh", "-c", cmdstr)
		cmd.Start()
		cmd.Wait()

        return c.File("./" + tmpFileNmae)
    }
}

func main() {
	fmt.Printf("Start GO echo...")
	resBody := response{Message: "hello"}

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
		return c.JSON(http.StatusOK, resBody)
	})

	e.GET("/grass/:username", MainPage())
	
	e.Logger.Fatal(e.Start(":8181"))

}
