package stringutils

func Substr(str string, start, end int) string {
	if len(str) < end || len(str) < start || end < start {
		return str
	}

	return str[start:end]
}
