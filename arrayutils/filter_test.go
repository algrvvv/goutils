package arrayutils

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		input    []any
		callback func(elem any) bool
		expected []any
	}{
		{
			input: []any{"hello", "world", "some", "tests"},
			callback: func(elem any) bool {
				if elem != "hello" {
					return false
				}
				return true
			},
			expected: []any{"hello"},
		},
		{
			input: []any{1, 2, 3, 4, 5, 6, 7, 8, 9},
			callback: func(elem any) bool {
				if elem.(int)%2 == 0 {
					return true
				}
				return false
			},
			expected: []any{2, 4, 6, 8},
		},
	}

	for _, test := range tests {
		actual := Filter(test.input, test.callback)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Filter(%v): expected %v, actual %v", test.input, test.expected, actual)
		}
	}

}
