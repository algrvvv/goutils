package dateutils

import (
	"testing"
)

func TestChangeDateFormatDefault(t *testing.T) {
	tests := []struct {
		inputDate       string
		inputDateLayout string
		format          string
		expected        string
	}{
		{"2024-06-02", "2006-01-02", "02.01.2006", "02.06.2024"},
	}

	for _, test := range tests {
		actual, err := ChangeDateFormatDefault(test.inputDate, test.inputDateLayout, test.format)
		if err != nil {
			t.Error(err)
		}
		if actual != test.expected {
			t.Errorf("Input: %v, Expected: %v, Actual: %v", test.inputDate, test.expected, actual)
		}
	}
}
