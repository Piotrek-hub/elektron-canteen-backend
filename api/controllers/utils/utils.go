package utils

import (
	"strconv"
	"time"
)

// FORMAT CZASU rok-miesiac-dzien

func UnixToDate(unix string) time.Time {
	timestamp, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(timestamp, 0)
}

func UnixToFormattedDate(unix string) string {
	timestamp, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		panic(err)
	}
	t := time.Unix(timestamp, 0)

	return t.Format("2006-01-02")
}

func GetDatesBetween(first, second time.Time) []string {
	var res = []string{first.Format("2006-01-02")}
	diff := int(second.Sub(first).Hours() / 24)

	tmpDate := first
	for i := 0; i < diff; i++ {
		t := tmpDate.AddDate(0, 0, 1)
		res = append(res, t.Format("2006-01-02"))
		tmpDate = t
	}

	return res
}
