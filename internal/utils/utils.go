package utils

import "strconv"

// BoolsToStr returns a string representation of the provided bool values.
func BoolsToStr(in ...bool) string {
	var res string

	for i, b := range in {
		if i == 0 {
			res += strconv.FormatBool(b)
		} else {
			res += "_" + strconv.FormatBool(b)
		}
	}

	return res
}
