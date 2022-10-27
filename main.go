package main

import (
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// must register counter on init
	prometheus.MustRegister(getBookCounter)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", bookHandler)

	println("listening..")
	http.ListenAndServe(":5005", nil)
}

// create a new counter vector
var getBookCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_books_count", // metric name
		Help: "Number of get_books request.",
	},
	[]string{"status"}, // labels
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	var status string
	defer func() {
		// increment the counter on defer func
		getBookCounter.WithLabelValues(status).Inc()
	}()

	var allStatus []string = []string{"success", "error"}
	var min, max int = 0, 1
	var index = rand.Intn(max-min+1) + min

	status = allStatus[index]
	w.Write([]byte(status))
}
