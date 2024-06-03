package arrayutils

import "slices"

func Distinct[S ~[]E, E comparable](s S) []E {
	var dist []E
	for _, v := range s {
		if !slices.Contains(dist, v) {
			dist = append(dist, v)
		}
	}

	return dist
}
