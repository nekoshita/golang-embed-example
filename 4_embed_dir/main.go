package main

import (
	"embed"
	"fmt"
)

//go:embed sql/*
var sql embed.FS

func main() {
	sample1, err := sql.ReadFile("sql/sample1.sql")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(sample1))

	sample2, err := sql.ReadFile("sql/sample2.sql")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(sample2))
}
