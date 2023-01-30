package utils

import "strconv"

func StringToInt(s string) (int, error) {
	num, err := strconv.ParseUint(s, 10, 64)
	return int(num), err
}

func IntToString(i int) string {
	return strconv.FormatUint(uint64(i), 10)
}
