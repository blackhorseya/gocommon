package stringutil

import "github.com/blackhorseya/gocommon/pkg/utils/intersect"

// CountIntersectStringSlice count two string slice intersect number
func CountIntersectStringSlice(s1, s2 []string) int {
	return len(intersect.Simple(s1, s2))
}

// InterfaceSliceToStringSlice convert interface slice to string slice
func InterfaceSliceToStringSlice(params ...interface{}) []string {
	var ret []string
	for _, param := range params {
		s, ok := param.(string)
		if !ok {
			continue
		}

		ret = append(ret, s)
	}

	return ret
}
