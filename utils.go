package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type result struct {
	Name     string     `json:"name"`
	Status   status     `json:"status,omitempty"`
	Perfdata []perfdata `json:"perfdata,omitempty"`
}

type Results struct {
	Results []result `json:"results"`
}

type perfdata struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type status struct {
	NumHostsUp   float64 `json:"num_hosts_up"`
	NumHostsDown float64 `json:"num_hosts_down"`
	Uptime       float64 `json:"uptime"`
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func JoinPath(base string, elem string) string {
	b := strings.TrimLeft(base, "/")
	return b + elem
}

func getMetrics(url string) result {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+BasicAuth(apiUsername, apiPassword))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: apiInsecure},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var res Results
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return res.Results[0]
}
