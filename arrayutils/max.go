package arrayutils

// Max function to find the maximum number in an array
func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxValue := arr[0]
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
