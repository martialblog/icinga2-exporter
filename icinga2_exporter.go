package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var build = "development"

var (
	listenAddress = flag.String("web.listen-address", ":9665", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
	msqlMetrics   = flag.Bool("msql", false, "Enable MySQL Metrics")
	apiBaseURL    = "https://localhost:5665/v1"
	apiUsername   = "root"
	apiPassword   = "password"
	apiInsecure   = false
)

func init() {
	if baseurl, provided := os.LookupEnv("ICINGA2_EXPORTER_BASE_URL"); provided {
		apiBaseURL = baseurl
	}

	if username, provided := os.LookupEnv("ICINGA2_EXPORTER_USERNAME"); provided {
		apiUsername = username
	}

	if password, provided := os.LookupEnv("ICINGA2_EXPORTER_PASSWORD"); provided {
		apiPassword = password
	} else {
		log.Printf("Warning: No API Password provided")
	}

	if _, provided := os.LookupEnv("ICINGA2_EXPORTER_TLS_INSECURE"); provided {
		apiInsecure = true
	}

	// Register Collectors
	prometheus.MustRegister(NewIcinga2CIBCollector())
	prometheus.MustRegister(NewIcinga2APICollector())
}

func main() {
	flag.Parse()

	if *msqlMetrics {
		prometheus.MustRegister(NewIcinga2MySQLCollector())
	}

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>Icinga2 Exporter</title></head>
			<body>
			<h1>Icinga2 Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Printf("Version: %s", build)
	log.Printf("Listening on address: %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
