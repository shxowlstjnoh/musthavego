package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("len(Str) = %d\n", len(str))
	fmt.Printf("len(runes) = %d\n", len(runes))
	
}
