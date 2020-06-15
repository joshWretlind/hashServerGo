package main

import (
	"encoding/json"
	"io/ioutil"
	"model"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var HASHED_PASSWORD = "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
func runServer() {
	go RunApplication()
}

func shutdownServer() {
	http.PostForm("http://localhost:12345/shutdown", url.Values{})
}

func TestServerRun(t *testing.T) {
	runServer()
}

func TestHashPostHappy(t *testing.T) {
	var response,error = http.PostForm("http://localhost:12345/hash", url.Values{"password":{"angryMonkey"}})
	if error != nil {
		t.Errorf("Error making post call %s", error)
	}
	defer response.Body.Close()
	var body, readError = ioutil.ReadAll(response.Body)
	if readError != nil {
		t.Errorf("Error reading response body %s", readError)
	}
	if string(body) != "1" {
		t.Errorf("Expected 1 from first post call, got %s", body)
	}
}

func TestHashGetHappy(t *testing.T) {
	http.PostForm("http://localhost:12345/hash", url.Values{"password":{"angryMonkey"}})
	var response, error = http.Get("http://localhost:12345/hash/2")
	if error != nil {
		t.Errorf("Error making get hash call %s", error)
	}
	var body, readError = ioutil.ReadAll(response.Body)
	if readError != nil {
		t.Errorf("Error reading response body %s", readError)
	}
	if string(body) != "" {
		t.Errorf("Error fetching hash, should be empty, was %s", body)
	}
	response.Body.Close()

	time.Sleep(6*time.Second)

	response, error = http.Get("http://localhost:12345/hash/1")
	if error != nil {
		t.Errorf("Error making get call %s", error)
	}
	body, readError = ioutil.ReadAll(response.Body)
	if readError != nil {
		t.Errorf("Error reading response body %s", readError)
	}
	if string(body) != HASHED_PASSWORD {
		t.Errorf("Error getting password. Should have been %s but was %s", HASHED_PASSWORD, body)
	}
	response.Body.Close()
}

func TestStatsGet(t *testing.T) {
	var response, error = http.Get("http://localhost:12345/stats")
	if error != nil {
		t.Errorf("Error making get statistics call %s", error)
	}
	var body, readError = ioutil.ReadAll(response.Body)
	if readError != nil {
		t.Errorf("Error reading response body %s", readError)
	}
	if string(body) == "" {
		t.Errorf("Error getting statistics body, was empty")
	}

	var stats model.Statistics
	var jsonError = json.Unmarshal(body, &stats)
	if jsonError != nil {
		t.Errorf("Error Unmarshalling statistics %s", jsonError)
	}
	if stats.TotalRequests != 2 {
		t.Errorf("More calls to server than was expected, expected 2, got %d", stats.TotalRequests)
	}
	if stats.AverageRequestTime < 5000.0 {
		t.Errorf("Average Request time was less than 5 seconds, should always be more than 5 seconds. Actual Request time: %f", stats.AverageRequestTime)
	}
}

func TestServerShutdown(t *testing.T) {
	shutdownServer()
}