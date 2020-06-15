package dao

import (
	"model"
	"sync"
)

var totalRequests uint64 = 0
var totalTime float64 = 0
var statsLock = &sync.Mutex{}

func AddRequestTime(requestTime float64) {
	statsLock.Lock()
	totalRequests++
	totalTime += requestTime
	statsLock.Unlock()
}

func GetStatistics() model.Statistics {
	return model.Statistics{totalRequests,totalTime/float64(totalRequests)}
}