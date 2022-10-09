package util

import "math/rand"

// [from, to)
func RandInt(start, stop int) int {
	return rand.Intn(stop-start) + start
}
