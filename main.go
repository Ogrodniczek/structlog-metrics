package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func recordMetrics() {
	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				var json_data interface{}
				data := fmt.Sprint(scanner.Text())
				_ = json.Unmarshal([]byte(data), &json_data)
				pat, _ := jsonpath.Compile(`$.status`)
				res, _ := pat.Lookup(json_data)
				fmt.Println(res)
				opsProcessed.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
			}

			if err := scanner.Err(); err != nil {
				log.Println(err)
			}
		}
	}()
}

var (
	opsProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mmmmmmmmmmmmmmmmmmmmmyapp_processed_ops_total",
			Help: "The total number of processed events",
		},
		[]string{"device"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(opsProcessed)
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}
