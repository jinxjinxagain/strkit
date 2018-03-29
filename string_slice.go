package strkit

// Map iterates through strings
// Return the strings after processed by proc func
func Map(s []string, proc func(ch string) string) []string {
	var r = []string{}
	for _, c := range s {
		var ac = proc(c)
		r = append(r, ac)
	}
	return r
}
