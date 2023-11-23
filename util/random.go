package util

import "math/rand"

func RandomInt(start, end int) int {
	return start + rand.Intn(end-start+1)
}
