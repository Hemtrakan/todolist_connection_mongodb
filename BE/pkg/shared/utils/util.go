package utils

import (
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func GetErrorcode(skip int) (errorCode string) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		errorCode = "SYSTEM_ERROR_"
		return errorCode
	}
	funcName := runtime.FuncForPC(pc).Name()
	idx := strings.LastIndex(funcName, "/")
	errorCode = "ERROR_" + strings.Replace(strings.Replace(strings.ToUpper(funcName[idx+1:]), ".(*", "_", -1), ").", "_", -1) + "_"
	return errorCode
}

func IntUnique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func StringUnique(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func IndexOfArray(element string, data []string) int {
	for k, regx := range data {
		match, _ := regexp.MatchString(regx, element)
		if match {
			return k
		}
	}
	return -1
}

func RemoveValueOfArray[T comparable](values []T, item T) []T {
	for i, other := range values {
		if other == item {
			return append(values[:i], values[i+1:]...)
		}
	}
	return values
}

func FilterVowels(input string) string {
	// สร้างฟังก์ชันเพื่อกรองสระออกจากสตริง
	filtered := strings.Map(func(r rune) rune {
		switch r {
		case 'ิ', 'ี', 'ึ', 'ื', 'ุ', 'ู', '็', '่', '้', '๊', '๋':
			// กรองสระออก
			return -1
		default:
			return r
		}
	}, input)

	return filtered
}

func AddComma(num int) string {
	// ตรวจสอบว่าตัวเลขเป็นค่าลบหรือไม่
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num // แปลงให้เป็นบวกเพื่อให้ได้ตัวเลขที่ถูกต้อง
	}

	// ใช้ strconv.Itoa() เพื่อแปลงจำนวนเต็มเป็น string
	numStr := strconv.Itoa(num)

	// หาความยาวของ string เพื่อใช้ในการวนลูป
	length := len(numStr)

	// หากความยาวของ string น้อยกว่าหรือเท่ากับ 3 ไม่ต้องทำการใส่ comma
	if length <= 3 {
		if isNegative {
			numStr = "-" + numStr
		}
		return numStr
	}

	// สร้าง string builder เพื่อการทำงานกับ string อย่างมีประสิทธิภาพ
	var formattedNum string

	// หา index ที่จะเริ่มใส่ comma
	start := length % 3

	// ถ้า start เท่ากับ 0 ให้ start เป็น 3 เพื่อให้เริ่มจากตัวที่ 3
	if start == 0 {
		start = 3
	}

	// วนลูปเพื่อใส่ comma ทุก 3 ตัว
	for i := 0; i < length; i++ {
		if i > 0 && (i-start)%3 == 0 {
			formattedNum += ","
		}
		formattedNum += string(numStr[i])
	}

	// ถ้าค่าเป็นลบให้เติมเครื่องหมายลบด้านหน้า
	if isNegative {
		formattedNum = "-" + formattedNum
	}

	return formattedNum
}
