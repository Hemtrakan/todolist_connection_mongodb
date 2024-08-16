package utils

import (
	"fmt"
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/shared/constant"
	"strings"
	"time"
)

func GetDate(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func GetDateFiscalYear(strDate string) (string, string) {

	dateConvert, _ := time.Parse(string(constant.DATE_FORMAT_STR), strDate)
	var dateFr = ""
	var dateTo = ""
	if dateConvert.Month() > 9 {
		firstDateOfYear := time.Date(dateConvert.Year()+1, time.January, 1, 0, 0, 0, 0, dateConvert.Location())
		dateFr = firstDateOfYear.AddDate(-1, -3, 0).Format(string(constant.DATE_FORMAT_STR))
		dateTo = strDate

	} else {

		firstDateOfYear := time.Date(dateConvert.Year(), time.January, 1, 0, 0, 0, 0, dateConvert.Location())
		dateFr = firstDateOfYear.AddDate(0, -3, 0).Format(string(constant.DATE_FORMAT_STR))
		dateTo = strDate
	}
	return dateFr, dateTo
}

func SetFileNameFormatPDF(Name string) (FileName string) {
	t := time.Now()
	formatYear := t.Format(string(constant.DATE_FORMAT_YYYYMMDD))
	formatTime := t.Format(string(constant.DATETIME_FORMAT))
	TimeSplit := strings.Split(formatTime, "T")
	getTimes := TimeSplit[1]
	getTimes = strings.Replace(getTimes, ":", "", 2)
	FileName = fmt.Sprintf("%v_%v_%v", Name, formatYear, getTimes)
	return
}

const (
	Layout = "02/01/2006"
)

func FormatDateString(t time.Time, layout string) string {
	return t.Format(layout)
}

func ConvertTimeFormat(t time.Time) (res string) {
	res = ""
	if !isTimeNil(t) {
		t, _ = time.Parse(Layout, t.Format(Layout))
		res = t.Format(Layout)
	}

	return
}

func ConvertTimeSetFormat(t time.Time, layout string) (res string) {
	res = ""
	if !isTimeNil(t) {
		t, _ = time.Parse(layout, t.Format(layout))
		res = t.Format(layout)
	}

	return
}

func ConvertStringToTime(t string) (res time.Time) {
	if t == "" {
		return
	}

	layout := "2006-01-02T15:04:05Z"
	res, err := time.Parse(layout, t)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	return
}

func isTimeNil(t time.Time) bool {
	return t.IsZero()
}

func ConvertStringToTimeLayout(strTime, layout string) (res time.Time, err error) {
	res, err = time.Parse(layout, strTime)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	return
}

func BetweenDate(before, after, main time.Time) (res bool) {
	res = false
	if main.Before(before) || main.After(after) {
		res = true
	}

	return
}
