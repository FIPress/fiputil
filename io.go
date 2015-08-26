package fiputil

import "bytes"

type StringWriter struct {
	buf bytes.Buffer
}

func (sw *StringWriter) Write(p []byte) (n int, err error) {
	return sw.buf.Write(p)
}

func (sw *StringWriter) String() string {
	return sw.buf.String()
}
