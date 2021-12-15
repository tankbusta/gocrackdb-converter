package utils

import (
	"strconv"
	"strings"
)

// IntSliceToString converts a list of integers to a comma separated list
func IntSliceToString(ints []int) string {
	tmp := make([]string, len(ints))
	for i, v := range ints {
		tmp[i] = strconv.Itoa(v)
	}

	return strings.Join(tmp, ",")
}
