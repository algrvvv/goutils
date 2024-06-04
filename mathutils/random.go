package mathutils

import "math/rand"

// RandomInt function for generating a random number in the desired range
func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
