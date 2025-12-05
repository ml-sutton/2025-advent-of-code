package utils

import "strconv"

func ParseInt(numberAsString string) int64 {
	parsedInt, err := strconv.ParseInt(numberAsString, 10, 64)
	if err != nil {
		return -1
	}
	return parsedInt
}
