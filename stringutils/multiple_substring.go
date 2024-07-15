package stringutils

import "strings"

// CheckSubstring a function that returns true only
// if all occurrences are in the past string
func CheckSubstring(s string, subs ...string) bool {
	var (
		countOfSubstrings    = len(subs)
		countOfAvailableSubs = 0
	)
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			countOfAvailableSubs++
		}
	}

	return countOfAvailableSubs == countOfSubstrings
}
