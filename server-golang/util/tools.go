package util

import "strconv"

func String2Int64(s string) (int64, error) {
	number, err := strconv.ParseInt(s, 10, 64)
	return number, err
}
