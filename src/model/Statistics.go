package model

type Statistics struct {
	TotalRequests uint64 "json:totalRequests"
	AverageRequestTime float64 "json:averageRequestTime"
}