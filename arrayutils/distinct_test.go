package arrayutils

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	tests := []struct {
		inputData []any
		expected  []any
	}{
		{[]any{1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9, 10}, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]any{9, 4, 5, -4, 9, 3, 6, 4}, []any{9, 4, 5, -4, 3, 6}},
		{[]any{"abc", "cba", "abc", "abc"}, []any{"abc", "cba"}},
		{[]any{'a', 'b', 'c', 'c'}, []any{'a', 'b', 'c'}},
	}

	for _, test := range tests {
		actual := Distinct(test.inputData)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Distinct(%+v): expected %v, actual %v", test.inputData, test.expected, actual)
		}
	}
}
