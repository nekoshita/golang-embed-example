package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed todo.json user.json
var responseJson embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(responseJson)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
