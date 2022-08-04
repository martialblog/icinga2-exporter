package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
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

func getMetrics(url string) result {
	var res Results

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic "+BasicAuth(apiUsername, apiPassword))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: apiInsecure},
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	resp, httpErr := client.Do(req)

	if httpErr != nil {
		log.Print("Error: ", httpErr)
		return result{}
	}

	if resp.StatusCode != http.StatusOK {
		log.Print("Error: ", resp.StatusCode, " ", url)
		return result{}
	}

	defer resp.Body.Close()
	jsonErr := json.NewDecoder(resp.Body).Decode(&res)

	// Since the status Object in the API has nested Objects
	// Let's ignore that for now
	if _, ok := jsonErr.(*json.UnmarshalTypeError); !ok && jsonErr != nil {
		log.Print("Error:", jsonErr)
		return result{}
	}

	// TODO There's gotta be a better way?
	if len(res.Results) > 0 {
		return res.Results[0]
	}

	return result{}
}
