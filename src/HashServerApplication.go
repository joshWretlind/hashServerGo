package main

import (
	"log"
	"net/http"
	"endpoints"
	"utils"
)

const NUM_WORKERS = 5
func main() {
	RunApplication()
}

func RunApplication() {
	utils.StartUpHashingWorkers(NUM_WORKERS)
	http.HandleFunc("/hash", endpoints.HashEndpoint)
	http.HandleFunc("/hash/", endpoints.HashEndpoint)
	http.HandleFunc("/stats", endpoints.StatsEndpoint)
	http.HandleFunc("/shutdown", endpoints.ShutdownEndpoint)
	var httpServer = http.Server{
		Addr: ":12345",
		Handler: nil,
	}
	utils.StartUpShutdownWorkers(httpServer)
	log.Fatal(httpServer.ListenAndServe())
}