package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func grafana_ok(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func grafana_metrics(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func grafana_metric_payload_options(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func grafana_query(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/grafana", grafana_ok)
	http.HandleFunc("/grafana/metrics", grafana_metrics)
	http.HandleFunc("/grafana/metric-payload-options", grafana_metric_payload_options)
	http.HandleFunc("/grafana/query", grafana_query)

	http.ListenAndServe(":8090", nil)
}
