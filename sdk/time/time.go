package time

import (
	"errors"
	"time"
)

func ToTimestamp(timeStr string) (timestamp int64, err error) {
	if len(timeStr) == 0 {
		timestamp = time.Now().Unix()
		return
	}
	parse, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		return
	}
	timestamp = parse.Unix()
	return
}

func ToDate(timestamp int64) (date string, err error) {
	if timestamp == -1 {
		date = time.Now().Local().Format(time.DateTime)
		return
	}
	if timestamp < 0 {
		err = errors.New("timestamp cannot be negative")
		return
	}
	date = time.Unix(timestamp, 0).Local().Format(time.DateTime)
	return
}
