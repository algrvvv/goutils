package extra

import "testing"

func TestTernary(t *testing.T) {
	tests := []struct {
		condition  bool
		trueValue  interface{}
		falseValue interface{}
		expected   interface{}
	}{
		{1+5 == 6, "This is true!", false, "This is true!"},
		{"test" == "test", 10, 41, 10},
		{"test" == "t3st", "Something wrong!", "Strings arent eq", "Strings arent eq"},
	}

	for _, test := range tests {
		actual := Ternary(test.condition, test.trueValue, test.falseValue)
		if actual != test.expected {
			t.Errorf("Ternary(%v, %v, %v) = %v; expected %v", test.condition, test.trueValue, test.falseValue, actual, test.expected)
		}
	}

}
