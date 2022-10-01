package main

import (
	"log"
	"net/http"
)

func main() {

	ServeStatic()
	StartServer("8000")

}

// Listen and serve port
func StartServer(addr string) {
	log.Printf("Starting server listening on: http://localhost:%v", addr)
	log.Fatal(http.ListenAndServe(":"+addr, nil))
}

// Gets the relative path of the requested path, making `website` the base path
func RelPath(path string) string {
	return "website" + path
}

// Provides handler function to read and respond with static files
func ServeStatic() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		path := RelPath(r.URL.Path)
		http.ServeFile(w, r, path)
	}
	http.HandleFunc("/", h1)
}
