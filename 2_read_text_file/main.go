package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed sample_bytes.txt sample_string.txt
var static embed.FS

func main() {
	b, err := static.ReadFile("sample_bytes.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(b))

	b2, err := fs.ReadFile(static, "sample_string.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(b2))
}
