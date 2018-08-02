package strkit

// Map iterates through strings
// Return the strings after processed by proc func
func SliceMap(s []string, proc func(ch string) string) []string {
	var r = []string{}
	for _, c := range s {
		var ac = proc(c)
		r = append(r, ac)
	}
	return r
}

// InSlice checks whether a target string in slice
func InSlice(s []string, t string) bool {
	for _, ss := range s {
		if t == ss {
			return true
		}
	}
	return false
}

// InSortedSlice checks whether a target string in sorted slice
// faster than normal InSlice
func InSortedSlice(s []string, t string) bool {
	var (
		low  = 0
		high = len(s) - 1
	)
	for low < high {
		var mid = (low + high) >> 1
		if s[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return s[low] == t
}

// Interfaces transfer slice of strings to slice of interfaces
func Interfaces(s []string) []interface{} {
	var r = []interface{}{}
	for _, ss := range s {
		r = append(r, ss)
	}
	return r
}
