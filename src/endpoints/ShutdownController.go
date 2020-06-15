package endpoints

import (
	"net/http"
	"utils"
)

var IsShuttingDown = false

func ShutdownEndpoint(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		shutdown(response, request)
	default:
		http.Error(response, "Unable to find handler for request", 405)
	}
}

func shutdown(response http.ResponseWriter, request *http.Request) {
	IsShuttingDown = true
	utils.ShutdownQueue <- 1
}