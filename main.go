package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	reqDur := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "",
			Name:      "request_duration_seconds",
			Help:      "request latencies",
			Buckets:   []float64{.005, .01, .02, 0.04, .06, 0.08, .1, 0.15, .25, 0.4, .6, .8, 1, 1.5, 2, 3, 5},
		},
		[]string{"code"},
	)

	prometheus.MustRegister(reqDur)

	router := echo.New()
	router.GET("/",
		func(ctx echo.Context) error {
			status := strconv.Itoa(http.StatusOK)
			elapsed := float64(time.Since(time.Now())) / float64(time.Second)
			reqDur.WithLabelValues(status).Observe(elapsed)
			i := time.Duration(rand.Intn(500))
			time.Sleep(i * time.Millisecond)
			return ctx.String(http.StatusOK, "server is working")
		},
	)
	router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	log.Fatal(router.Start(":8080"))
}
