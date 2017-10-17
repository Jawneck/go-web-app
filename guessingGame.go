//A web application in Go.
//Author: Danielis Joniskis

package main

import (
	"fmt"
	"net/http"
)

//A http.ResponseWriter assembles the HTTP server's response by writing to it.
//A http.Request is a data structure that represents the client HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Guessing game")
}

//The main function begins with a call to http.HandleFunc.
//It then calls http.ListenAndServe, specifying that it should listen on port 8080.
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}