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
[{
  "label": "Describe metric list",
  "value": "DescribeMetricList",
  "payloads": [{
    "label": "Namespace",
    "name": "namespace",
    "type": "select",
    "defaultValue": "acs_mongodb",
    "placeholder": "Please select namespace",
    "reloadMetric": true,
    "options": [{
      "label": "acs_mongodb",
      "value": "acs_mongodb"
    },{
      "label": "acs_rds",
      "value": "acs_rds"
    }]
  },{
    "name": "metric",
    "type": "select"
  }]
},{
  "value": "DescribeMetricLast",
  "payloads": [{
    "name": "namespace",
    "type": "select"
  },{
    "name": "metric",
    "type": "input"
  },{
    "name": "instanceId",
    "type": "multi-select"
  }]
}]
`
var rdsMetrics = `
[{
  "label": "Describe metric list",
  "value": "DescribeMetricList",
  "payloads": [{
    "label": "Namespace",
    "name": "namespace",
    "type": "select",
    "defaultValue": "acs_mongodb",
    "placeholder": "Please select namespace",
    "reloadMetric": true,
    "options": [{
      "label": "acs_mongodb",
      "value": "acs_mongodb"
    },{
      "label": "acs_rds",
      "value": "acs_rds"
    }]
  },{
    "name": "metric",
    "type": "select"
  },{
    "name": "instanceId",
    "type": "select"
  }]
},{
  "value": "DescribeMetricLast",
  "payloads": [{
    "name": "namespace",
    "type": "select"
  },{
    "name": "metric",
    "type": "input"
  },{
    "name": "instanceId",
    "type": "multi-select"
  }]
}]
`

var testData = `[
	{
	  "id": 0,
	  "time": 1698853664,
	  "value": 13
	},
	{
	  "id": 1,
	  "time": 1698853964,
	  "value": 35
	},
	{
	  "id": 2,
	  "time": 1698854264,
	  "value": 86
	},
	{
	  "id": 3,
	  "time": 1698854564,
	  "value": 96
	},
	{
	  "id": 4,
	  "time": 1698854864,
	  "value": 86
	},
	{
	  "id": 5,
	  "time": 1698855164,
	  "value": 71
	},
	{
	  "id": 6,
	  "time": 1698855464,
	  "value": 95
	},
	{
	  "id": 7,
	  "time": 1698855764,
	  "value": 26
	},
	{
	  "id": 8,
	  "time": 1698856064,
	  "value": 56
	},
	{
	  "id": 9,
	  "time": 1698856364,
	  "value": 44
	},
	{
	  "id": 10,
	  "time": 1698856664,
	  "value": 97
	},
	{
	  "id": 11,
	  "time": 1698856964,
	  "value": 42
	},
	{
	  "id": 12,
	  "time": 1698857264,
	  "value": 46
	},
	{
	  "id": 13,
	  "time": 1698857564,
	  "value": 12
	},
	{
	  "id": 14,
	  "time": 1698857864,
	  "value": 98
	},
	{
	  "id": 15,
	  "time": 1698858164,
	  "value": 66
	},
	{
	  "id": 16,
	  "time": 1698858464,
	  "value": 23
	},
	{
	  "id": 17,
	  "time": 1698858764,
	  "value": 75
	},
	{
	  "id": 18,
	  "time": 1698859064,
	  "value": 30
	},
	{
	  "id": 19,
	  "time": 1698859364,
	  "value": 31
	},
	{
	  "id": 20,
	  "time": 1698859664,
	  "value": 94
	},
	{
	  "id": 21,
	  "time": 1698859964,
	  "value": 3
	},
	{
	  "id": 22,
	  "time": 1698860264,
	  "value": 2
	},
	{
	  "id": 23,
	  "time": 1698860564,
	  "value": 56
	},
	{
	  "id": 24,
	  "time": 1698860864,
	  "value": 36
	},
	{
	  "id": 25,
	  "time": 1698861164,
	  "value": 88
	},
	{
	  "id": 26,
	  "time": 1698861464,
	  "value": 61
	},
	{
	  "id": 27,
	  "time": 1698861764,
	  "value": 87
	},
	{
	  "id": 28,
	  "time": 1698862064,
	  "value": 38
	},
	{
	  "id": 29,
	  "time": 1698862364,
	  "value": 96
	},
	{
	  "id": 30,
	  "time": 1698862664,
	  "value": 93
	},
	{
	  "id": 31,
	  "time": 1698862964,
	  "value": 82
	},
	{
	  "id": 32,
	  "time": 1698863264,
	  "value": 77
	},
	{
	  "id": 33,
	  "time": 1698863564,
	  "value": 27
	},
	{
	  "id": 34,
	  "time": 1698863864,
	  "value": 94
	},
	{
	  "id": 35,
	  "time": 1698864164,
	  "value": 4
	},
	{
	  "id": 36,
	  "time": 1698864464,
	  "value": 62
	},
	{
	  "id": 37,
	  "time": 1698864764,
	  "value": 62
	},
	{
	  "id": 38,
	  "time": 1698865064,
	  "value": 20
	},
	{
	  "id": 39,
	  "time": 1698865364,
	  "value": 43
	},
	{
	  "id": 40,
	  "time": 1698865664,
	  "value": 80
	},
	{
	  "id": 41,
	  "time": 1698865964,
	  "value": 54
	},
	{
	  "id": 42,
	  "time": 1698866264,
	  "value": 52
	},
	{
	  "id": 43,
	  "time": 1698866564,
	  "value": 81
	},
	{
	  "id": 44,
	  "time": 1698866864,
	  "value": 28
	},
	{
	  "id": 45,
	  "time": 1698867164,
	  "value": 31
	},
	{
	  "id": 46,
	  "time": 1698867464,
	  "value": 14
	},
	{
	  "id": 47,
	  "time": 1698867764,
	  "value": 52
	},
	{
	  "id": 48,
	  "time": 1698868064,
	  "value": 35
	},
	{
	  "id": 49,
	  "time": 1698868364,
	  "value": 50
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
		writer.Write([]byte(rdsMetrics))
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
	w.Write([]byte(testData))
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
