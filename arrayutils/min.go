package arrayutils

// MinInt function to find the minimum number in an array
func MinInt(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	minInt := arr[0]
	for _, v := range arr {
		if v < minInt {
			minInt = v
		}
	}
	return minInt
}
