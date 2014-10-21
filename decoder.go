package jsonmin

import (
	"encoding/json"
	"io"
)

// NewDecoder() returns a new json.Decoder that reads from r
// and allows comments, making it suitable for config files.
func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(NewReader(r))
}

// Unmarshal parses the JSON-encoded data (possibly with comments)
// and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	min := minifier{state: startState}
	n := min.Bytes(data)
	return json.Unmarshal(data[:n], v)
}
