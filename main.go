package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	hitsTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hits_total",
	})
	if err := prometheus.Register(hitsTotal); err != nil {
		log.Fatal(err)
	}
	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		log.Fatal(err)
	}
	//go func() {
	//	metricsRouter := echo.New()
	//	log.Fatal(metricsRouter.Start(":5050"))
	//}()

	router := echo.New()
	router.GET("/",
		func(ctx echo.Context) error {
			i := time.Duration(rand.Intn(500))
			time.Sleep(i * time.Millisecond)
			return ctx.String(http.StatusOK, "server is working")
		},
	)
	router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	log.Fatal(router.Start(":8080"))
}
