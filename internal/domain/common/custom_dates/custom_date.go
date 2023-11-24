package custom_dates

import "time"

func TodayBeginningHour() time.Time {
	return time.Date(time.Now().Local().Year(), time.Now().Local().Month(), time.Now().Local().Day(), 0, 0, 0, 0, time.Local)
}
