package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type Message struct{
    Msg string `json :"text"`
}

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) (err error) {
        return c.String(http.StatusOK, "Hello, World!")
    })

	e.POST("/post" , func(c echo.Context) error {
        var message Message
        if err :=c.Bind(&message); err!= nil{
            return c.JSON(http.StatusBadRequest ,map[string]string{"error":"invaled values or parameters"})
        }
        return c.JSON(http.StatusOK , map[string]string{"received":message.Msg})
        
    })
    e.Logger.Fatal(e.Start(":1323"))
}