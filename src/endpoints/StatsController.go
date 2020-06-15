package endpoints

import (
	"encoding/json"
	"net/http"
	"fmt"
	"dao"
)

func StatsEndpoint(response http.ResponseWriter, request *http.Request) {
	if IsShuttingDown {
		http.Error(response, "Service is shutting down, retry later", 503)
	}
	//This bugs me as far as a code-standards point of view is concerned.
	//I'd rather be able to specify which methods handle which verbs in the server config.
	//Similar to how the gorilla lib works, but not being able to use 3rd party libraries here
	switch request.Method {
	case http.MethodGet:
		statsGet(response, request)
	default:
		http.Error(response, "Unable to find handler for request", 405)
	}
}
func statsGet(response http.ResponseWriter, request *http.Request) {
	var stats = dao.GetStatistics()
	var results,_ = json.Marshal(stats)
	fmt.Fprint(response, string(results))
}
