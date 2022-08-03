package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type icinga2APICollector struct {
	apinumhttpclients *prometheus.Desc
}

func NewIcinga2APICollector() *icinga2APICollector {
	return &icinga2APICollector{
		apinumhttpclients: prometheus.NewDesc("icinga2_api_numhttpclients", "Number of HTTP Clients", nil, nil),
	}
}

func (collector *icinga2APICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.apinumhttpclients
}

func (collector *icinga2APICollector) Collect(ch chan<- prometheus.Metric) {
	// TOOD Golang 1.19 https://pkg.go.dev/net/url@master#JoinPath
	url := JoinPath(apiBaseURL, "/status/ApiListener")
	icinga := getMetrics(url)

	// Transform to map so that we can access it easily
	var perfdata = make(map[string]float64)
	for _, v := range icinga.Perfdata {
		perfdata[v.Label] = v.Value
	}

	ch <- prometheus.MustNewConstMetric(collector.apinumhttpclients, prometheus.GaugeValue, perfdata["api_num_http_clients"])
}
