package multiply

import (
	"testing"
)

func TestMultiply(t *testing.T) {

	var tests = []struct {
		num1 string
		num2 string
		want string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"0", "9", "0"},
		{"498828660196", "840477629533", "419254329864656431168468"},
		{"1011", "6", "6066"},
	}
	for _, tt := range tests {
		got := multiply(tt.num1, tt.num2)
		if got != tt.want {
			t.Errorf("%s * %s expected: %v; got: %v", tt.num1, tt.num2, tt.want, got)
		}
	}
}
