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
	Name     string             `json:"name"`
	Status   map[string]float64 `json:"status,omitempty"`
	Perfdata []perfdata         `json:"perfdata,omitempty"`
}

type Results struct {
	Results []result `json:"results"`
}

type perfdata struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
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

	resp, httpErr := client.Do(req)

	if httpErr != nil {
		log.Fatal(httpErr)
	}

	var res Results
	jsonErr := json.NewDecoder(resp.Body).Decode(&res)

	// Since the status Object in the API has nested Objects
	// Let's ignore that for now
	if jsonErr != nil && jsonErr != jsonErr.(*json.UnmarshalTypeError) {
		log.Fatal(jsonErr)
	}

	defer resp.Body.Close()
	return res.Results[0]
}
