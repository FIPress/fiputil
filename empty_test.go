package fiputil

import (
	"testing"
	"unsafe"
)

var e3 Empty

func TestEmpty(t *testing.T) {
	e := Empty{}
	t.Log("size:", unsafe.Sizeof(e))
	t.Logf("address: %p\n", &e)

	var e1 Empty
	t.Log("size:", unsafe.Sizeof(e1))
	t.Logf("address: %p\n", &e1)

	t.Log("size:", unsafe.Sizeof(e3))
	t.Logf("address: %p\n", &e3)

}
