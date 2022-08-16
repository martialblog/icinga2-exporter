package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/url"
)

type icinga2APICollector struct {
	api_num_conn_endpoints                 *prometheus.Desc
	api_num_not_conn_endpoints             *prometheus.Desc
	api_num_endpoints                      *prometheus.Desc
	api_num_http_clients                   *prometheus.Desc
	api_num_json_rpc_anonymous_clients     *prometheus.Desc
	api_num_json_rpc_relay_queue_item_rate *prometheus.Desc
	api_num_json_rpc_relay_queue_items     *prometheus.Desc
	api_num_json_rpc_sync_queue_item_rate  *prometheus.Desc
	api_num_json_rpc_sync_queue_items      *prometheus.Desc
	api_num_json_rpc_work_queue_item_rate  *prometheus.Desc
}

func NewIcinga2APICollector() *icinga2APICollector {
	return &icinga2APICollector{
		api_num_conn_endpoints:                 prometheus.NewDesc("icinga2_api_num_conn_endpoints", "Number of connected Endpoints", nil, nil),
		api_num_endpoints:                      prometheus.NewDesc("icinga2_api_num_endpoints", "Number of Endpoints", nil, nil),
		api_num_http_clients:                   prometheus.NewDesc("icinga2_api_num_http_clients", "Number of HTTP Clients", nil, nil),
		api_num_json_rpc_anonymous_clients:     prometheus.NewDesc("icinga2_api_num_json_rpc_anonymous_clients", "Number of JSON RPC Anonymous Clients", nil, nil),
		api_num_json_rpc_relay_queue_item_rate: prometheus.NewDesc("icinga2_api_num_json_rpc_relay_queue_item_rate", "JSON RPC relay queue item rate", nil, nil),
		api_num_json_rpc_relay_queue_items:     prometheus.NewDesc("icinga2_api_num_json_rpc_relay_queue_items", "JSON RPC relay queue items", nil, nil),
		api_num_json_rpc_sync_queue_item_rate:  prometheus.NewDesc("icinga2_api_num_json_rpc_sync_queue_item_rate", "JSON RPC sync queue item rate", nil, nil),
		api_num_json_rpc_sync_queue_items:      prometheus.NewDesc("icinga2_api_num_json_rpc_sync_queue_items", "JSON RPC sync queue items", nil, nil),
		api_num_json_rpc_work_queue_item_rate:  prometheus.NewDesc("icinga2_api_num_json_rpc_work_queue_item_rate", "JSON RPC work queue item rate", nil, nil),
		api_num_not_conn_endpoints:             prometheus.NewDesc("icinga2_api_num_not_conn_endpoints", "Number of not connected Endpoints", nil, nil),
	}
}

func (collector *icinga2APICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.api_num_conn_endpoints
	ch <- collector.api_num_endpoints
	ch <- collector.api_num_http_clients
	ch <- collector.api_num_json_rpc_anonymous_clients
	ch <- collector.api_num_json_rpc_relay_queue_item_rate
	ch <- collector.api_num_json_rpc_relay_queue_items
	ch <- collector.api_num_json_rpc_sync_queue_item_rate
	ch <- collector.api_num_json_rpc_sync_queue_items
	ch <- collector.api_num_json_rpc_work_queue_item_rate
	ch <- collector.api_num_not_conn_endpoints
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
	ch <- prometheus.MustNewConstMetric(collector.api_num_endpoints, prometheus.GaugeValue, perfdata["api_num_endpoints"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_http_clients, prometheus.GaugeValue, perfdata["api_num_http_clients"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_json_rpc_anonymous_clients, prometheus.GaugeValue, perfdata["api_num_json_rpc_anonymous_clients"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_json_rpc_relay_queue_item_rate, prometheus.GaugeValue, perfdata["api_num_json_rpc_relay_queue_item_rate"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_json_rpc_relay_queue_items, prometheus.GaugeValue, perfdata["api_num_json_rpc_relay_queue_items"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_json_rpc_sync_queue_item_rate, prometheus.GaugeValue, perfdata["api_num_json_rpc_sync_queue_item_rate"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_json_rpc_sync_queue_items, prometheus.GaugeValue, perfdata["api_num_json_rpc_sync_queue_items"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_json_rpc_work_queue_item_rate, prometheus.GaugeValue, perfdata["api_num_json_rpc_work_queue_item_rate"])
	ch <- prometheus.MustNewConstMetric(collector.api_num_not_conn_endpoints, prometheus.GaugeValue, perfdata["api_num_not_conn_endpoints"])
}
