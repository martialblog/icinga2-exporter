package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/url"
)

type icinga2CIBCollector struct {
	uptime             *prometheus.Desc
	num_hosts_up       *prometheus.Desc
	num_hosts_down     *prometheus.Desc
	avg_execution_time *prometheus.Desc
	avg_latency        *prometheus.Desc
	max_execution_time *prometheus.Desc
	max_latency        *prometheus.Desc
	min_execution_time *prometheus.Desc
	min_latency        *prometheus.Desc
}

func NewIcinga2CIBCollector() *icinga2CIBCollector {
	return &icinga2CIBCollector{
		uptime:             prometheus.NewDesc("icinga2_uptime", "Uptime of the instance", nil, nil),
		num_hosts_up:       prometheus.NewDesc("icinga2_num_hosts_up", "Number of Hosts Up", nil, nil),
		num_hosts_down:     prometheus.NewDesc("icinga2_num_hosts_down", "Number of Hosts Down", nil, nil),
		avg_execution_time: prometheus.NewDesc("icinga2_avg_execution_time", "Average execution time", nil, nil),
		avg_latency:        prometheus.NewDesc("icinga2_avg_latency", "Average latency", nil, nil),
		max_execution_time: prometheus.NewDesc("icinga2_max_execution_time", "Maximum execution time", nil, nil),
		max_latency:        prometheus.NewDesc("icinga2_max_latency", "Maximum latency", nil, nil),
		min_execution_time: prometheus.NewDesc("icinga2_min_execution_time", "Minimum execution time", nil, nil),
		min_latency:        prometheus.NewDesc("icinga2_min_latency", "Minimum latency", nil, nil),
	}
}

func (collector *icinga2CIBCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.uptime
	ch <- collector.num_hosts_up
	ch <- collector.num_hosts_down
	ch <- collector.avg_execution_time
	ch <- collector.avg_latency
	ch <- collector.max_execution_time
	ch <- collector.max_latency
	ch <- collector.min_execution_time
	ch <- collector.min_latency
}

func (collector *icinga2CIBCollector) Collect(ch chan<- prometheus.Metric) {
	url, _ := url.JoinPath(apiBaseURL, "/status/CIB")
	icinga := getMetrics(url).Status

	ch <- prometheus.MustNewConstMetric(collector.uptime, prometheus.GaugeValue, icinga["uptime"])
	ch <- prometheus.MustNewConstMetric(collector.num_hosts_up, prometheus.GaugeValue, icinga["num_hosts_up"])
	ch <- prometheus.MustNewConstMetric(collector.num_hosts_down, prometheus.GaugeValue, icinga["num_hosts_down"])
	ch <- prometheus.MustNewConstMetric(collector.avg_execution_time, prometheus.GaugeValue, icinga["avg_execution_time"])
	ch <- prometheus.MustNewConstMetric(collector.avg_latency, prometheus.GaugeValue, icinga["avg_latency"])
	ch <- prometheus.MustNewConstMetric(collector.max_execution_time, prometheus.GaugeValue, icinga["max_execution_time"])
	ch <- prometheus.MustNewConstMetric(collector.max_latency, prometheus.GaugeValue, icinga["max_latency"])
	ch <- prometheus.MustNewConstMetric(collector.min_execution_time, prometheus.GaugeValue, icinga["min_execution_time"])
	ch <- prometheus.MustNewConstMetric(collector.min_latency, prometheus.GaugeValue, icinga["min_latency"])
}
