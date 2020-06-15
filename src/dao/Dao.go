package dao

import (
	"sync"
)

var size = 0
var hashMap = make(map[int]string)
var dbLock = &sync.Mutex{}

//"Database" like operations
func CreateRecord(record string) int {
	dbLock.Lock()
	size++
	hashMap[size] = record
	dbLock.Unlock()
	return size
}

func GetRecord(index int) string {
	return hashMap[index]
}

func UpdateRecord(index int, value string) {
	dbLock.Lock()
	hashMap[index] = value
	dbLock.Unlock()
}