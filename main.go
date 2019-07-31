package main

import (
        "net/http"
        "time"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
        go func() {
                for {
                        myGague.Add(11)
                        time.Sleep(2 * time.Second)
                }
        }()
}

var (
         myGague = promauto.NewGauge(prometheus.GaugeOpts{
         Name: "my_example_gauge_data",
         Help: "my example gauge data",
         ConstLabels:map[string]string{"error":""},
         })
)

func main() {
        recordMetrics()

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}
