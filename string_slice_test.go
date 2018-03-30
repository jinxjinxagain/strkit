package strkit

import "testing"

var AppendXCharFunc = func(x string) string {
	return x + "x"
}

func TestMapMethod(t *testing.T) {
	var st = []string{"1", "2", "3", "4", "5"}
	var tt = Map(st, AppendXCharFunc)
	for index := range st {
		expected(t, st[index]+"x" == tt[index], st[index]+"x", tt[index])
	}
}

func TestInSlice(t *testing.T) {
	var s = []string{"3", "2", "1", "4", "5", "6"}
	var ok1 = InSlice(s, "4")
	expected(t, ok1, true, ok1)
	var ok2 = InSlice(s, "7")
	expected(t, !ok2, false, ok2)
}

func TestInSortedSlice(t *testing.T) {
	var s = []string{"1", "2", "3", "4", "5", "6"}
	var ok1 = InSortedSlice(s, "3")
	expected(t, ok1, true, ok1)
	var ok2 = InSortedSlice(s, "7")
	expected(t, !ok2, false, ok2)
}

func expected(t testing.TB, cond bool, expected, got interface{}) {
	if cond {
		return
	}
	t.Fatalf("expected '%v', but got '%v'", expected, got)
}
