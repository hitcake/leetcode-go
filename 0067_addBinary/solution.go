package addBinary

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	var buf strings.Builder
	var carry int = 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		a1 := 0
		if i >= 0 {
			a1 = int(a[i] - '0')
		}
		b1 := 0
		if j >= 0 {
			b1 = int(b[j] - '0')
		}
		cur := a1 + b1 + carry
		if cur >= 2 {
			carry = 1
			cur -= 2
		} else {
			carry = 0
		}
		buf.WriteString(strconv.Itoa(cur))
	}
	s := []rune(buf.String())
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
