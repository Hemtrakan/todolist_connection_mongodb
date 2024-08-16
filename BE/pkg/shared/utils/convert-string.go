package utils

import (
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/shared/constant"
	"strconv"
	"time"
)

func ConvertStringToDate(dateStr string) time.Time {
	val, _ := time.Parse(string(constant.DATETIME_FORMAT), dateStr)
	return val
}

func ConvertStringToFloat(floatStr string) float64 {
	val, _ := strconv.ParseFloat(floatStr, 64)
	return val
}

func ConvertToDate(dateStr string) time.Time {
	val, _ := time.Parse(string(constant.DATE_FORMAT_STR), dateStr)
	return val
}

func ConvertToDateTime(dateStr string) time.Time {

	inputLayout := "02-01-2006 15:04:05"

	parsedTime, _ := time.Parse(inputLayout, dateStr)

	outputLayout := "2006-01-02 15:04:05.0000000"

	output := parsedTime.Format(outputLayout)
	val, _ := time.Parse(string(constant.DATETIME_FORMAT_STR2), output)
	return val
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func ConvertStringToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return res
}

func ConvertDatetoString(date time.Time) string {
	formattedTime := date.Format(string(constant.DATE_FORMAT_STR))
	return formattedTime
}
