package arrayutils

// Filter a function that takes an array and filters it by callback
func Filter[T any](arr []T, fn func(T) bool) []T {
	var newArr []T
	for _, v := range arr {
		if fn(v) {
			newArr = append(newArr, v)
		}
	}

	return newArr
}
