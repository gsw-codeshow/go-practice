package task

import (
	"time"
)

var timeNow = time.Now

func GetUnixTime() int64 {
	return timeNow().Unix()
}
