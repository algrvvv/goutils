package arrayutils

// MapValues a function to iterate through an array of elements
// and apply a callback to each element
func MapValues[T any](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}

	return newValues
}
