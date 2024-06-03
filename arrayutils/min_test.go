package arrayutils

import "testing"

func TestMinInt(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 5, 3, 8, 3}, 1},
		{[]int{}, 0},
		{[]int{94, 43, 7, 11, 87, 2, 02, -1, 4, 92}, -1},
	}

	for _, test := range tests {
		actual := MinInt(test.numbers)
		if actual != test.expected {
			t.Errorf("MinInt(%v) = %v; expected %v", test.numbers, actual, test.expected)
		}
	}
}
