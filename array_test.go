package fiputil

import "testing"

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	arrMaped := Map(arr, func(a interface{}) interface{} {
		i, _ := a.(int)
		return i + 1
	})
	t.Log(arrMaped)
}

func TestMkString(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Log(MkString(arr, "(", ",", ")"))
}
