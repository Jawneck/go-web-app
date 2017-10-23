//A web application in Go.
//Author: Danielis Joniskis

package main

import (
	"fmt"
	"net/http"
    "html/template"
    "os"
)

//A http.ResponseWriter assembles the HTTP server's response by writing to it.
//A http.Request is a data structure that represents the client HTTP request.
/*
func handler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type","text/html")
    fmt.Fprintln(w, "<h1>Guessing game</h1>")
}
*/

type Template struct {
    Message string
}

//The main function begins with a call to http.HandleFunc.
//It then calls http.ListenAndServe, specifying that it should listen on port 8080.
func main() {
    //http.HandleFunc("/", handler)
    http.Handle("/",  http.FileServer(http.Dir("./")))//serving index.html in the root directory.
    http.ListenAndServe(":8080", nil)

     t := template.New("Template")
     t, _ = t.Parse("<h2>Guess a number between 1 and 20</h2> {{.Message}}!")
     t.Execute(os.Stdout)


}