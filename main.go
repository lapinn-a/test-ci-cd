package main

import (
	"encoding/json"
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

var defaultMetrics = `
[
  {
    "id": 0,
    "time": 1698848364,
    "value": 52
  },
  {
    "id": 1,
    "time": 1698848664,
    "value": 14
  },
  {
    "id": 2,
    "time": 1698848964,
    "value": 100
  },
  {
    "id": 3,
    "time": 1698849264,
    "value": 34
  },
  {
    "id": 4,
    "time": 1698849564,
    "value": 68
  },
  {
    "id": 5,
    "time": 1698849864,
    "value": 80
  },
  {
    "id": 6,
    "time": 1698850164,
    "value": 27
  },
  {
    "id": 7,
    "time": 1698850464,
    "value": 7
  },
  {
    "id": 8,
    "time": 1698850764,
    "value": 26
  },
  {
    "id": 9,
    "time": 1698851064,
    "value": 95
  },
  {
    "id": 10,
    "time": 1698851364,
    "value": 68
  },
  {
    "id": 11,
    "time": 1698851664,
    "value": 9
  },
  {
    "id": 12,
    "time": 1698851964,
    "value": 59
  },
  {
    "id": 13,
    "time": 1698852264,
    "value": 93
  },
  {
    "id": 14,
    "time": 1698852564,
    "value": 22
  },
  {
    "id": 15,
    "time": 1698852864,
    "value": 66
  },
  {
    "id": 16,
    "time": 1698853164,
    "value": 2
  },
  {
    "id": 17,
    "time": 1698853464,
    "value": 84
  },
  {
    "id": 18,
    "time": 1698853764,
    "value": 50
  },
  {
    "id": 19,
    "time": 1698854064,
    "value": 70
  },
  {
    "id": 20,
    "time": 1698854364,
    "value": 99
  },
  {
    "id": 21,
    "time": 1698854664,
    "value": 3
  },
  {
    "id": 22,
    "time": 1698854964,
    "value": 98
  },
  {
    "id": 23,
    "time": 1698855264,
    "value": 95
  },
  {
    "id": 24,
    "time": 1698855564,
    "value": 76
  },
  {
    "id": 25,
    "time": 1698855864,
    "value": 15
  },
  {
    "id": 26,
    "time": 1698856164,
    "value": 71
  },
  {
    "id": 27,
    "time": 1698856464,
    "value": 6
  },
  {
    "id": 28,
    "time": 1698856764,
    "value": 71
  },
  {
    "id": 29,
    "time": 1698857064,
    "value": 63
  },
  {
    "id": 30,
    "time": 1698857364,
    "value": 18
  },
  {
    "id": 31,
    "time": 1698857664,
    "value": 44
  },
  {
    "id": 32,
    "time": 1698857964,
    "value": 59
  },
  {
    "id": 33,
    "time": 1698858264,
    "value": 27
  },
  {
    "id": 34,
    "time": 1698858564,
    "value": 75
  },
  {
    "id": 35,
    "time": 1698858864,
    "value": 26
  },
  {
    "id": 36,
    "time": 1698859164,
    "value": 66
  },
  {
    "id": 37,
    "time": 1698859464,
    "value": 82
  },
  {
    "id": 38,
    "time": 1698859764,
    "value": 49
  },
  {
    "id": 39,
    "time": 1698860064,
    "value": 87
  },
  {
    "id": 40,
    "time": 1698860364,
    "value": 84
  },
  {
    "id": 41,
    "time": 1698860664,
    "value": 74
  },
  {
    "id": 42,
    "time": 1698860964,
    "value": 84
  },
  {
    "id": 43,
    "time": 1698861264,
    "value": 18
  },
  {
    "id": 44,
    "time": 1698861564,
    "value": 14
  },
  {
    "id": 45,
    "time": 1698861864,
    "value": 17
  },
  {
    "id": 46,
    "time": 1698862164,
    "value": 64
  },
  {
    "id": 47,
    "time": 1698862464,
    "value": 81
  },
  {
    "id": 48,
    "time": 1698862764,
    "value": 34
  },
  {
    "id": 49,
    "time": 1698863064,
    "value": 100
  }
]
`

type MetricsRequest struct {
	Metric  string                 `json:"metric"`
	Payload map[string]interface{} `json:"payload"`
}

func grafana_ok(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func grafana_metrics(writer http.ResponseWriter, request *http.Request) {
	var req MetricsRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	writer.Header().Set("content-type", "application/json")
	if req.Metric == "DescribeMetricList" && req.Payload["namespace"] == "acs_rds" {
		writer.Write([]byte("test"))
		return
	}
	writer.Write([]byte(defaultMetrics))
}

type OptionsRequest struct {
	Name    string                 `json:"name"`
	Metric  string                 `json:"metric"`
	Payload map[string]interface{} `json:"payload"`
}

func grafana_metric_payload_options(writer http.ResponseWriter, request *http.Request) {
	var req OptionsRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	writer.Header().Set("content-type", "application/json")
	switch req.Name {
	case "instanceId":
		writer.Write([]byte(`[{
      "label": "My Database 1",
      "value": "sadbip2kasdmnlo"
    },{
      "label": "My Database 2",
      "value": "sadbip2kasdmnla"
    },{
      "label": "My Database 3",
      "value": "sadbip2kasdmnlx"
    }]`))

	case "metric":
		writer.Write([]byte(`[{
      "label": "CPUUtilization",
      "value": "CPUUtilization"
    },{
      "label": "DiskReadIOPS",
      "value": "DiskReadIOPS"
    },{
      "label": "memory_freeutilization",
      "value": "memory_freeutilization"
    }]`))
	case "namespace":
		writer.Write([]byte(`[{
      "label": "MongoDB",
      "value": "acs_mongodb"
    },{
      "label": "RDS",
      "value": "acs_rds"
    },{
      "label": "Load balancer",
      "value": "acs_load_balancer"
    }]`))
	}
}

func grafana_query(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(defaultMetrics))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/grafana/", grafana_ok)
	http.HandleFunc("/grafana/metrics", grafana_metrics)
	http.HandleFunc("/grafana/metric-payload-options", grafana_metric_payload_options)
	http.HandleFunc("/grafana/query", grafana_query)

	http.ListenAndServe(":8090", nil)
}
