package fiputil

import "testing"

func TestSlug(t *testing.T) {
	in := "**我&是 标abc -- -题 &&"
	out := Slug(in)
	if out != "我-是-标abc-题" {
		t.Error("get slug failed")
	}
}
