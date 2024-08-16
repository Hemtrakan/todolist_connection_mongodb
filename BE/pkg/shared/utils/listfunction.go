package utils

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func UniqueStr(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		_, value := keys[entry]
		if !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		_, value := keys[entry]
		if !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
