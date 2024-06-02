package stringutils

import (
	"strings"
)

// CompareWithoutCase function that checks string similarity in a case-insensitive manner
func CompareWithoutCase(firstStr, secondStr string) bool {
	return strings.ToLower(firstStr) == strings.ToLower(secondStr)
}

// CompareWithCase function that checks string similarity, taking into account case
func CompareWithCase(firstStr, secondStr string) bool {
	return firstStr == secondStr
}
