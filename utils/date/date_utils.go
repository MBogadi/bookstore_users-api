package date

import "time"

const dateFormat = "2006-01-02T15:04:05-0700" // this string has to be 2006-01... based string!!!  or you can use RFC3339

func GetNowUTC() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	//return GetNowUTC().Format(time.RFC3339)
	return GetNowUTC().Format("2006-01-02 15:04:05")
}
