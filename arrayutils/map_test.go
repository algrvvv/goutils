package arrayutils

import (
	"reflect"
	"strings"
	"testing"
)

func TestMapValues(t *testing.T) {
	tests := []struct {
		input    []interface{}
		expected []interface{}
		callback func(v interface{}) interface{}
	}{
		{[]interface{}{1, 2, 3}, []interface{}{2, 4, 6}, func(v interface{}) interface{} {
			return v.(int) * 2
		}},
		{[]interface{}{1.4, 2.8, 3.9}, []interface{}{4.34, 8.68, 12.09}, func(v interface{}) interface{} {
			return v.(float64) * 3.1
		}},
		{[]interface{}{"abc", "aaa", "bbc"}, []interface{}{"abcabcabc", "aaaaaaaaa", "bbcbbcbbc"}, func(v interface{}) interface{} {
			return strings.Repeat(v.(string), 3)
		}},
	}

	for _, test := range tests {
		actual := MapValues(test.input, test.callback)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("MapValues: expected %v, actual %v", test.expected, actual)
		}
	}
}
