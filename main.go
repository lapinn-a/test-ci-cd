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

var testData = `
[
  {
    "id": 0,
    "time": 1697587001621,
    "value": 0
  },
  {
    "id": 1,
    "time": 1696625856575,
    "value": 59
  },
  {
    "id": 2,
    "time": 1696648462607,
    "value": 25
  },
  {
    "id": 3,
    "time": 1698201058942,
    "value": 6
  },
  {
    "id": 4,
    "time": 1697101766834,
    "value": 70
  },
  {
    "id": 5,
    "time": 1698009496208,
    "value": 83
  },
  {
    "id": 6,
    "time": 1698713810804,
    "value": 40
  },
  {
    "id": 7,
    "time": 1698305894895,
    "value": 68
  },
  {
    "id": 8,
    "time": 1697995376724,
    "value": 90
  },
  {
    "id": 9,
    "time": 1697619803999,
    "value": 32
  },
  {
    "id": 10,
    "time": 1697073531675,
    "value": 60
  },
  {
    "id": 11,
    "time": 1698440982957,
    "value": 90
  },
  {
    "id": 12,
    "time": 1698031123421,
    "value": 60
  },
  {
    "id": 13,
    "time": 1698501079394,
    "value": 56
  },
  {
    "id": 14,
    "time": 1696669114269,
    "value": 21
  },
  {
    "id": 15,
    "time": 1696695469328,
    "value": 95
  },
  {
    "id": 16,
    "time": 1697348502489,
    "value": 0
  },
  {
    "id": 17,
    "time": 1696225399912,
    "value": 44
  },
  {
    "id": 18,
    "time": 1698142965727,
    "value": 61
  },
  {
    "id": 19,
    "time": 1696816556593,
    "value": 13
  },
  {
    "id": 20,
    "time": 1698118794143,
    "value": 97
  },
  {
    "id": 21,
    "time": 1696872373285,
    "value": 91
  },
  {
    "id": 22,
    "time": 1697311269457,
    "value": 75
  },
  {
    "id": 23,
    "time": 1698340203498,
    "value": 72
  },
  {
    "id": 24,
    "time": 1698713993433,
    "value": 9
  },
  {
    "id": 25,
    "time": 1696766109169,
    "value": 29
  },
  {
    "id": 26,
    "time": 1698050549096,
    "value": 6
  },
  {
    "id": 27,
    "time": 1696910646573,
    "value": 84
  },
  {
    "id": 28,
    "time": 1696695497257,
    "value": 90
  },
  {
    "id": 29,
    "time": 1698314601850,
    "value": 61
  },
  {
    "id": 30,
    "time": 1698826325180,
    "value": 77
  },
  {
    "id": 31,
    "time": 1697021820042,
    "value": 85
  },
  {
    "id": 32,
    "time": 1697762659642,
    "value": 33
  },
  {
    "id": 33,
    "time": 1696367078973,
    "value": 16
  },
  {
    "id": 34,
    "time": 1696494667425,
    "value": 18
  },
  {
    "id": 35,
    "time": 1698658671406,
    "value": 21
  },
  {
    "id": 36,
    "time": 1697374983131,
    "value": 1
  },
  {
    "id": 37,
    "time": 1696849345173,
    "value": 79
  },
  {
    "id": 38,
    "time": 1698050536794,
    "value": 88
  },
  {
    "id": 39,
    "time": 1697038309645,
    "value": 52
  },
  {
    "id": 40,
    "time": 1698465944412,
    "value": 29
  },
  {
    "id": 41,
    "time": 1697562268336,
    "value": 95
  },
  {
    "id": 42,
    "time": 1698395783780,
    "value": 45
  },
  {
    "id": 43,
    "time": 1696572624145,
    "value": 27
  },
  {
    "id": 44,
    "time": 1697000571838,
    "value": 6
  },
  {
    "id": 45,
    "time": 1696945871002,
    "value": 5
  },
  {
    "id": 46,
    "time": 1696926825656,
    "value": 34
  },
  {
    "id": 47,
    "time": 1698392858167,
    "value": 12
  },
  {
    "id": 48,
    "time": 1698412371378,
    "value": 94
  },
  {
    "id": 49,
    "time": 1696580337752,
    "value": 3
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
