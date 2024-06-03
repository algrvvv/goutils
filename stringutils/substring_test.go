package stringutils

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		str        string
		start, end int
		expected   string
	}{
		{"hello", 0, 2, "he"},
		{"hello", 2, 4, "ll"},
		{"hello", 8, 10, "hello"},
		{"hello", 10, 2, "hello"},
	}

	for _, test := range tests {
		actual := Substr(test.str, test.start, test.end)
		if actual != test.expected {
			t.Errorf("\nExpected: %v\nActual: %v", test.expected, actual)
		}
	}

}
