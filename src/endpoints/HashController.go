package endpoints

import (
	"dao"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"utils"
	"model"
)

const PASSWORD_FIELD = "password"

func HashEndpoint(response http.ResponseWriter, request *http.Request) {
	if IsShuttingDown {
		http.Error(response, "Service is shutting down, retry later", 503)
	}
	//This bugs me as far as a code-standards point of view is concerned.
	//I'd rather be able to specify which methods handle which verbs in the server config.
	//Similar to how the gorilla lib works, but not being able to use 3rd party libraries here
	switch request.Method {
	case http.MethodGet:
		hashGet(response, request)
	case http.MethodPost:
		hashPost(response, request)
	default:
		http.Error(response, "Unable to find handler for request", 405)
	}
}

func hashPost(response http.ResponseWriter, request *http.Request) {
	var beginTime = uint64(time.Now().UnixNano())
	request.ParseForm()
	var password = request.Form.Get(PASSWORD_FIELD)
	var index = dao.CreateRecord("")
	var hashRecord = model.HashRecord{index,password,beginTime}
	utils.JobWaitGroup.Add(1)
	utils.HashingJobQueue <- hashRecord
	fmt.Fprint(response, index)
}

func hashGet(response http.ResponseWriter, request *http.Request) {
	var id, error = strconv.Atoi(strings.TrimPrefix(request.URL.Path, "/hash/"))
	if error != nil {
		http.Error(response, "There was an error processing your request", 400)
		return
	}
	fmt.Fprint(response, dao.GetRecord(id))
}