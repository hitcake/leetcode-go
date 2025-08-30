package addBinary

import "testing"

func TestAddBinary(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want string
	}{
		{"11", "1", "100"},
		{"1011", "1010", "10101"},
		{"1111", "1111", "11110"},
	}
	for _, v := range tests {
		if got := addBinary(v.a, v.b); got != v.want {
			t.Errorf("%v and %v should be %v, got %v", v.a, v.b, v.want, got)
		}
	}
}
