package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample_bytes.txt
var sampleBytes []byte

//go:embed sample_string.txt
var sampleString string

func main() {
	fmt.Printf("%s\n", sampleBytes)
	fmt.Printf("%s\n", sampleString)
}
