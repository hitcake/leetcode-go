package cashier

import (
	"testing"
)

func TestCashier_GetBill(t *testing.T) {
	cashier := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	var tests = []struct {
		products []int
		prices   []int
		want     float64
	}{
		{[]int{1, 2}, []int{1, 2}, 500.0},
		{[]int{3, 7}, []int{10, 10}, 4000.0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1}, 800.0},
		{[]int{4}, []int{10}, 4000.0},
		{[]int{7, 3}, []int{10, 10}, 4000.0},
		{[]int{7, 5, 3, 1, 6, 4, 2}, []int{10, 10, 10, 9, 9, 9, 7}, 7350.0},
		{[]int{2, 3, 5}, []int{5, 3, 2}, 2500.0},
	}
	for _, test := range tests {
		got := cashier.GetBill(test.products, test.prices)
		if got != test.want {
			t.Errorf("GetBill(%v, %v) = %v, want %v", test.products, test.prices, got, test.want)
		}
	}
}
