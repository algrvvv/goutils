package extra

// Ternary a function that simulates the behavior of the ternary operator
func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}

	return falseVal
}
