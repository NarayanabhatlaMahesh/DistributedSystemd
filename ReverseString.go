package main

import "fmt"

func main() {
	s := "Hello, World!"
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	fmt.Println(string(b))
}