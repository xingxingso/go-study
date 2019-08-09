package main

import "golang.org/x/tour/reader"

type MyReader struct {}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
    // b = b[:cap(b)]
    for i := range b {
        b[i] = 65
        // b[i] = 'A'
    }
    // return cap(b), nil
    return len(b), nil
}

func main() {
    reader.Validate(MyReader{})
}
