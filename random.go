package fiputil

import "math/rand"

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Random8() uint8 {
	return uint8(rand.Int63() >> 55)
}
