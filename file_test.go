package fiputil

import "testing"

func TestGetExtension(t *testing.T) {
	path := "abd.cf"
	ext := GetExtension(path)
	if ext != "cf" {
		t.Log("ext:", ext)
		t.Fail()
	}
}

func TestRemoveExtension(t *testing.T) {
	path := "abd.cf"
	name := RemoveExtension(path)
	if name != "abd" {
		t.Log("name:", name)
		t.Fail()
	}
}
