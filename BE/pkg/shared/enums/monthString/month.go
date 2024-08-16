package monthString

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (s Month) String() string {
	return [...]string{
		"มกราคม",
		"กุมภาพันธ์",
		"มีนาคม",
		"เมษายน",
		"พฤษภาคม",
		"มิถุนายน",
		"กรกฎาคม",
		"สิงหาคม",
		"กันยายน",
		"ตุลาคม",
		"พฤศจิกายน",
		"ธันวาคม"}[s-1]
}

func GetMonthInt(numberOfMonth int) (res string) {
	switch numberOfMonth {
	case 1:
		res = "มกราคม"
		return
	case 2:
		res = "กุมภาพันธ์"
		return
	case 3:
		res = "มีนาคม"
		return
	case 4:
		res = "เมษายน"
		return
	case 5:
		res = "พฤษภาคม"
		return
	case 6:
		res = "มิถุนายน"
		return
	case 7:
		res = "กรกฎาคม"
		return
	case 8:
		res = "สิงหาคม"
		return
	case 9:
		res = "กันยายน"
		return
	case 10:
		res = "ตุลาคม"
		return
	case 11:
		res = "พฤศจิกายน"
		return
	case 12:
		res = "ธันวาคม"
		return
	default:
		res = "ไม่เจอเดือนที่ต้องการ"
		return
	}
	 
}
