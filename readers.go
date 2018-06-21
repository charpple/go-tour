package main

import "golang.org/x/tour/reader"
//import "strings"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (int, error) {
        n := cap(b)
        for i := 0; i<n; i++ {
                b[i] = 'A'
        }
        //b = strings.Repeat('A', n)
        return n, nil
}


func main() {
        reader.Validate(MyReader{})
}
