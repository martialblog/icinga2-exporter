package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	expected := "Zm9vYmFyOmZvb2Jhcg=="
	actual := BasicAuth("foobar", "foobar")

	if expected != actual {
		t.Fatalf(`actual == %s ; expected == %s`, actual, expected)
	}
}

type MetricTest struct {
	name          string
	server        *httptest.Server
	response      result
	expectedError error
}

func TestGetMetrics(t *testing.T) {
	tests := []MetricTest{
		{
			name: "cib-ok",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"results":[{"name":"CIB","perfdata":[],"status":{"active_host_checks":0.1,"active_host_checks_15min":15}}]}`))
			})),
			response: result{
				Name:     "CIB",
				Status:   map[string]float64{"active_host_checks": 0.1, "active_host_checks_15min": 15},
				Perfdata: []perfdata{},
			},
			expectedError: nil,
		},
		{
			name: "api-ok",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"results":[{"name":"API","perfdata":[{"label":"api_endpoints", "value": 1.0}],"status":{}}]}`))
			})),
			response: result{
				Name:     "API",
				Status:   map[string]float64{},
				Perfdata: []perfdata{{Label: "api_endpoints", Value: 1.0}},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()
			actual := getMetrics(test.server.URL)
			if !reflect.DeepEqual(actual.Status, test.response.Status) || !reflect.DeepEqual(actual.Perfdata, test.response.Perfdata) {
				t.Fatal("actual:", actual, "expected:", test.response)
			}
		})
	}
}
