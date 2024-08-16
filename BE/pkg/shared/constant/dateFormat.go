package constant

type DateFormat string

const (
	DATE_FORMAT           DateFormat = "2006-01-02T00:00:00"
	DATETIME_FORMAT       DateFormat = "2006-01-02T15:04:05"
	DATEDOCUMENT_FORMAT   DateFormat = "060102"
	DATE_FORMAT_YYYYMMDD  DateFormat = "20060102"
	DATE_FORMAT_DISPLAY   DateFormat = "02/01/2006"
	DATE_FORMAT_STR       DateFormat = "2006-01-02"
	DATE_FORMAT_STR1      DateFormat = "02-01-2006"
	DATE_FORMAT_SQL       DateFormat = "2006-01-02T00:00:00Z"
	DATE_FORMAT_MONTHYEAR DateFormat = "01/2006"
	TIME_FORMAT           DateFormat = "15:04"
	DATETIME_FORMAT_STR2  DateFormat = "2006-01-02 15:04:05"
)
