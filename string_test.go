package fiputil

import (
	"fmt"
	"testing"
)

type Test struct {
	A string
}

type TestString struct {
	A string
}

func (ts TestString) String() string {
	return "test string " + ts.A
}

func TestToString(t *testing.T) {
	t.Log(ToString(12345))
	t.Log(ToString(12.345))
	t.Log(ToString(struct{}{}))
	t.Log(ToString(TestString{"anc"}))
	t.Log(fmt.Print(TestString{"anc"}))
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToString(12.345)
	}
}

func BenchmarkPrint(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprint(12.345)
	}
}
