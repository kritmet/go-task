package util

import "time"

// LoadLocation for get time zone
func LoadLocation() *time.Location {
	timeZone, _ := time.LoadLocation("Asia/Bangkok")
	return timeZone
}
