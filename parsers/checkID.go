package parsers

import (
	"strconv"
)

func CheckId(id string) bool {
	if id[0] == '0' {
		return false
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	if i >= 1 && i <= 52 {
		return true
	}
	return false
}
