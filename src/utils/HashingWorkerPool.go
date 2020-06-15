package utils

import (
	"dao"
	"sync"
	"time"
	"model"
	"crypto/sha512"
    "encoding/base64"
)

const SLEEP_TIME = 5 * time.Second
const QUEUE_SIZE = 100

//This is a buffered pool. If we run out of queue then this will block.
//Ideally we could split out the worker pool from the api/controller hosts, and be able to handle more throughput that way
var HashingJobQueue = make(chan model.HashRecord, QUEUE_SIZE)
var JobWaitGroup sync.WaitGroup

func StartUpHashingWorkers(numOfWorkers int) {
	for i := 1; i <= numOfWorkers; i++ {
		go hashWorker(HashingJobQueue)
	}
}

func hashWorker(jobQueue chan model.HashRecord) {
	for job := range jobQueue {
		time.Sleep(SLEEP_TIME)

		var sha512Hasher = sha512.New()
		sha512Hasher.Write([]byte(job.StringToHash))
		var hashedPassword = base64.StdEncoding.EncodeToString(sha512Hasher.Sum(nil))
		dao.UpdateRecord(job.Index, hashedPassword)

		var endTime = uint64(time.Now().UnixNano())
		var timeToAdd = float64(endTime - job.CreateTime)/float64(time.Millisecond)
		dao.AddRequestTime(timeToAdd)

		JobWaitGroup.Done()
	}
}

