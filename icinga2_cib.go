package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type icinga2CIBCollector struct {
	uptime       *prometheus.Desc
	numhostsup   *prometheus.Desc
	numhostsdown *prometheus.Desc
}

func NewIcinga2CIBCollector() *icinga2CIBCollector {
	return &icinga2CIBCollector{
		uptime:       prometheus.NewDesc("icinga2_uptime", "Uptime of the instance", nil, nil),
		numhostsup:   prometheus.NewDesc("icinga2_num_hosts_up", "Number of Hosts Up", nil, nil),
		numhostsdown: prometheus.NewDesc("icinga2_num_hosts_down", "Number of Hosts Down", nil, nil),
	}
}

func (collector *icinga2CIBCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.uptime
	ch <- collector.numhostsup
	ch <- collector.numhostsdown
}

func (collector *icinga2CIBCollector) Collect(ch chan<- prometheus.Metric) {
	// TOOD Golang 1.19 https://pkg.go.dev/net/url@master#JoinPath
	url := JoinPath(apiBaseURL, "/CIB")
	icinga := getMetrics(url).Status

	ch <- prometheus.MustNewConstMetric(collector.uptime, prometheus.GaugeValue, icinga["uptime"])
	ch <- prometheus.MustNewConstMetric(collector.numhostsup, prometheus.GaugeValue, icinga["num_hosts_up"])
	ch <- prometheus.MustNewConstMetric(collector.numhostsdown, prometheus.GaugeValue, icinga["num_hosts_down"])
}
