package jsonmin

import (
	"io"
)

type reader struct {
	r io.Reader
	c *minifier
}

// NewReader returns a reader that removes
// comments from JSON source. It does not perform
// any syntactic tests.
func NewReader(source io.Reader) io.Reader {
	return &reader{source, &minifier{state: startState}}
}

func (r *reader) Read(v []byte) (n int, err error) {
	nread := 0
	for n < len(v) && err == nil {
		nread, err = r.r.Read(v[n:])
		nread = r.c.Bytes(v[n : n+nread])
		n += nread
	}
	return
}
