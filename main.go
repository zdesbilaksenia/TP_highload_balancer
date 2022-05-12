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
	reqDur := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Subsystem: "",
			Name:      "request_duration_seconds",
			Help:      "request latencies",
			Buckets:   []float64{.005, .01, .02, 0.04, .06, 0.08, .1, 0.15, .25, 0.4, .6, .8, 1, 1.5, 2, 3, 5},
		},
	)

	prometheus.MustRegister(reqDur)

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
