package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/url"
)

type icinga2MySQLCollector struct {
	idomysqlconnection_ido_mysql_queries_rate          *prometheus.Desc
	idomysqlconnection_ido_mysql_queries_1min          *prometheus.Desc
	idomysqlconnection_ido_mysql_queries_5mins         *prometheus.Desc
	idomysqlconnection_ido_mysql_queries_15mins        *prometheus.Desc
	idomysqlconnection_ido_mysql_query_queue_items     *prometheus.Desc
	idomysqlconnection_ido_mysql_query_queue_item_rate *prometheus.Desc
}

func NewIcinga2MySQLCollector() *icinga2MySQLCollector {
	return &icinga2MySQLCollector{
		idomysqlconnection_ido_mysql_queries_rate:          prometheus.NewDesc("icinga_idomysqlconnection_ido_mysql_queries_rate", "MySQL queries rate", nil, nil),
		idomysqlconnection_ido_mysql_queries_1min:          prometheus.NewDesc("icinga_idomysqlconnection_ido_mysql_queries_1min", "MySQL queries 1 Minute", nil, nil),
		idomysqlconnection_ido_mysql_queries_5mins:         prometheus.NewDesc("icinga_idomysqlconnection_ido_mysql_queries_5mins", "MySQL queries 5 Minutes", nil, nil),
		idomysqlconnection_ido_mysql_queries_15mins:        prometheus.NewDesc("icinga_idomysqlconnection_ido_mysql_queries_15mins", "MySQL queries 15 Minutes", nil, nil),
		idomysqlconnection_ido_mysql_query_queue_items:     prometheus.NewDesc("icinga_idomysqlconnection_ido_mysql_query_queue_items", "MySQL query queue items", nil, nil),
		idomysqlconnection_ido_mysql_query_queue_item_rate: prometheus.NewDesc("icinga_idomysqlconnection_ido_mysql_query_queue_item_rate", "MySQL query queue rate", nil, nil),
	}
}

func (collector *icinga2MySQLCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.idomysqlconnection_ido_mysql_queries_rate
	ch <- collector.idomysqlconnection_ido_mysql_queries_rate
	ch <- collector.idomysqlconnection_ido_mysql_queries_1min
	ch <- collector.idomysqlconnection_ido_mysql_queries_5mins
	ch <- collector.idomysqlconnection_ido_mysql_queries_15mins
	ch <- collector.idomysqlconnection_ido_mysql_query_queue_items
	ch <- collector.idomysqlconnection_ido_mysql_query_queue_item_rate
}

func (collector *icinga2MySQLCollector) Collect(ch chan<- prometheus.Metric) {
	url, _ := url.JoinPath(apiBaseURL, "/status/IdoMysqlConnection")
	icinga := getMetrics(url)

	// Transform to map so that we can access it easily
	var perfdata = make(map[string]float64)
	for _, v := range icinga.Perfdata {
		perfdata[v.Label] = v.Value
	}

	ch <- prometheus.MustNewConstMetric(collector.idomysqlconnection_ido_mysql_queries_rate, prometheus.GaugeValue, perfdata["idomysqlconnection_ido-mysql_queries_rate"])
	ch <- prometheus.MustNewConstMetric(collector.idomysqlconnection_ido_mysql_queries_1min, prometheus.GaugeValue, perfdata["idomysqlconnection_ido-mysql_queries_1min"])
	ch <- prometheus.MustNewConstMetric(collector.idomysqlconnection_ido_mysql_queries_5mins, prometheus.GaugeValue, perfdata["idomysqlconnection_ido-mysql_queries_5mins"])
	ch <- prometheus.MustNewConstMetric(collector.idomysqlconnection_ido_mysql_queries_15mins, prometheus.GaugeValue, perfdata["idomysqlconnection_ido-mysql_queries_15mins"])
	ch <- prometheus.MustNewConstMetric(collector.idomysqlconnection_ido_mysql_query_queue_items, prometheus.GaugeValue, perfdata["idomysqlconnection_ido-mysql_query_queue_items"])
	ch <- prometheus.MustNewConstMetric(collector.idomysqlconnection_ido_mysql_query_queue_item_rate, prometheus.GaugeValue, perfdata["idomysqlconnection_ido-mysql_query_queue_item_rate"])
}
