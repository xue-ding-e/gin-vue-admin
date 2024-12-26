package utils

import "strconv"

// uint转化为string
func UintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
