package stringutils

import "testing"

func TestCompareWithoutCase(t *testing.T) {
	tests := []struct {
		firstStr  string
		secondStr string
		expected  bool
	}{
		{"Hello", "HeLLo", true},
		{"W0rld", "w0RLd", true},
		{"OlleH", "hElLo", false},
	}

	for _, test := range tests {
		actual := CompareWithoutCase(test.firstStr, test.secondStr)
		if actual != test.expected {
			t.Errorf("CompareWithoutCase(%q, %q) = %v; expected%v)", test.firstStr, test.secondStr, actual, test.expected)
		}
	}

}
