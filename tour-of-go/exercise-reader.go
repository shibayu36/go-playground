package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// func (r MyReader) Read(b []byte) (int, error) {
// 	l := len(b)
// 	for i := 0; i < l; i++ {
// 		b[i] = 'A'
// 	}
// 	return l, nil
// }

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	b := make([]byte, 8)
	(MyReader{}).Read(b)
	println(string(b))

	reader.Validate(MyReader{})
}
