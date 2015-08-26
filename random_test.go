package fiputil

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandom8(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(Random8())
		t.Log(rand.Int())
	}
	time.Sleep(5 * time.Microsecond)
	for i := 0; i < 5; i++ {
		t.Log(Random8())
		t.Log(rand.Int())
	}
}
