package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read method to implement io.Reader interface
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A' // fill buffer with character 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
