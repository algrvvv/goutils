package arrayutils

import "cmp"

func MapValues[T cmp.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}

	return newValues
}
