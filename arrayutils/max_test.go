package arrayutils

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 5, 3, 8, 3}, 8},
		{[]int{}, 0},
		{[]int{94, 43, 7, 11, 87, 2, 02, -1, 4, 92}, 94},
	}

	for _, test := range tests {
		actual := Max(test.numbers)
		if actual != test.expected {
			t.Errorf("Max(%#v) = %#v; expected %#v", test.numbers, actual, test.expected)
		}
	}
}
