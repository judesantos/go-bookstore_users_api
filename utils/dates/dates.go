package dates

import "time"

const (
	dateFormat = "YYYY-MM-DD HH:MM:SS"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(dateFormat)
}
