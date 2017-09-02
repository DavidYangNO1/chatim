package utility

import (
	"time"
)

func GetTimeStamp() int64 {
	timestamp := time.Now().Unix()
	return timestamp
}

func GetTimeStampWithTime(pointTime time.Time) int64 {
	return pointTime.Unix()
}

func GetTimeStampNano() float64 {
	timestampFloat := float64(time.Now().UnixNano()) / 1.0e9
	return timestampFloat
}

func ConverToTime(timestamp int64) time.Time {
	t1 := time.Unix(timestamp, 0)
	return t1
}

func ConvertToTimeByNano(nano float64) time.Time {
	t1 := time.Unix(int64(nano), int64((nano-float64(int64(nano)))*1e9))
	return t1
}
