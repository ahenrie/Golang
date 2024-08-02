// Front and back controller pattern
// Front controller (one object) received the request and routes the request to one
// or more back controllers.

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Serve up the menu
	// Store in file objectr
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
