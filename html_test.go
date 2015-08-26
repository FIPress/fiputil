package fiputil

import "testing"

func TestHtmlEscape(t *testing.T) {
	input := `abcd&cdef`
	output := HtmlEscape([]byte(input))
	s := string(output)
	if s != `abcd&amp;cdef` {
		t.Log(s)
		t.FailNow()
	}

	input = `abcd&cd<ef`
	output = HtmlEscape([]byte(input))
	s = string(output)
	if s != `abcd&amp;cd&lt;ef` {
		t.Log(s)
		t.FailNow()
	}

	input = `abcd&copy;d<ef`
	output = HtmlEscape([]byte(input))
	s = string(output)
	if s != `abcd&copy;d&lt;ef` {
		t.Log(s)
		t.FailNow()
	}
}

func TestHtmlEscapeLiterally(t *testing.T) {
	input := `abcd&cdef`
	output := HtmlEscapeLiterally([]byte(input))
	s := string(output)
	if s != `abcd&amp;cdef` {
		t.Log(s)
		t.FailNow()
	}

	input = `abcd&copy;d<ef`
	output = HtmlEscapeLiterally([]byte(input))
	s = string(output)
	if s != `abcd&amp;copy;d&lt;ef` {
		t.Log(s)
		t.FailNow()
	}
}
