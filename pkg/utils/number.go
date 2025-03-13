package utils

func IsNumber(str string) bool {
	//纯数字返回true
	for _, v := range str {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}