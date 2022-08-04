package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/url"
)

type icinga2APICollector struct {
	api_num_conn_endpoints     *prometheus.Desc
	api_num_not_conn_endpoints *prometheus.Desc
	api_num_endpoints          *prometheus.Desc
	api_num_http_clients       *prometheus.Desc
}

func NewIcinga2APICollector() *icinga2APICollector {
	return &icinga2APICollector{
		api_num_conn_endpoints:     prometheus.NewDesc("icinga2_api_num_conn_endpoints", "Number of connected Endpoints", nil, nil),
		api_num_not_conn_endpoints: prometheus.NewDesc("icinga2_api_num_not_conn_endpoints", "Number of not connected Endpoints", nil, nil),
		api_num_endpoints:          prometheus.NewDesc("icinga2_api_num_endpoints", "Number of Endpoints", nil, nil),
		api_num_http_clients:       prometheus.NewDesc("icinga2_api_num_http_clients", "Number of HTTP Clients", nil, nil),
	}
}

func (collector *icinga2APICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.api_num_conn_endpoints
	ch <- collector.api_num_not_conn_endpoints
	ch <- collector.api_num_endpoints
	ch <- collector.api_num_http_clients

}

func (collector *icinga2APICollector) Collect(ch chan<- prometheus.Metric) {

	url, _ := url.JoinPath(apiBaseURL, "/status/ApiListener")
	icinga := getMetrics(url)

	// Transform to map so that we can access it easily
	var perfdata = make(map[string]float64)
	for _, v := range icinga.Perfdata {
		perfdata[v.Label] = v.Value
	}

	ch <- prometheus.MustNewConstMetric(collector.api_num_conn_endpoints, prometheus.GaugeValue, perfdata["api_num_conn_endpoints"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_not_conn_endpoints, prometheus.GaugeValue, perfdata["api_num_not_conn_endpoints"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_endpoints, prometheus.GaugeValue, perfdata["api_num_endpoints"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_http_clients, prometheus.GaugeValue, perfdata["api_num_http_clients"])
}
