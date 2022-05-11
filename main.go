package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	router := echo.New()
	router.GET("/",
		func(ctx echo.Context) error {
			i := time.Duration(rand.Intn(1000))
			time.Sleep(i * time.Millisecond)
			return ctx.String(http.StatusOK, "server is working")
		},
	)
	log.Fatal(router.Start(":8080"))
}
