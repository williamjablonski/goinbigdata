package main

import (
	"encoding/json"
	"github.com/upitau/goinbigdata/examples/healthcheck/elastic"
	"github.com/upitau/goinbigdata/examples/healthcheck/health"
	"github.com/upitau/goinbigdata/examples/healthcheck/mongo"
	"net/http"
)

func main() {
	healthService := health.New([]string{"node1", "node2", "node3"}, mongo.New(), elastic.New())
	http.HandleFunc("/health", statusHandler(healthService))
	http.ListenAndServe("localhost:8080", nil)
}

func statusHandler(healthService health.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.MarshalIndent(healthService.Health(), "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}
