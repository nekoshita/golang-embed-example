package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var static embed.FS

func main() {
	public, err := fs.Sub(static, "static/public")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(public)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
