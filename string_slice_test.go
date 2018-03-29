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

func expected(t testing.TB, cond bool, expected, got interface{}) {
	if cond {
		return
	}
	t.Fatalf("expected '%v', but got '%v'", expected, got)
}
