//A web application in Go.
//Author: Danielis Joniskis

package main

import (
//   "fmt"
	"net/http"
    "text/template"
    "strconv"
)

//A http.ResponseWriter assembles the HTTP server's response by writing to it.
//A http.Request is a data structure that represents the client HTTP request.
/*func handler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type","text/html")
    fmt.Fprintln(w, "<h1>Guessing game</h1>")
}*/

type Template struct {
    Message string
    Count int
}

//Adapted from https://www.youtube.com/watch?v=GTSq1VPPFco&feature=youtu.be.
func templateHandler(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("template/guess.html")
    t.Execute(w, Template{Message: "Guess a number between 1 and 20"})

        i := 0
    
        //Reading the cookie //Adapted from https://github.com/data-representation/go-cookies
        var target, err = r.Cookie("i")
        if err == nil {
            // converting the value to an int
            i, _ = strconv.Atoi(target.Value)
        }
        i += 1

    //Creating and setting a session cookie
	target = &http.Cookie{
		Name:    "i",
		Value:   strconv.Itoa(i),
	}
	http.SetCookie(w, target)
}

//The main function begins with a call to http.Handle.
//It then calls http.ListenAndServe, specifying that it should listen on port 8080.
func main() {
    http.Handle("/",  http.FileServer(http.Dir("./static")))//serving index.html in the root directory.
    http.HandleFunc("/guess", templateHandler)
    http.ListenAndServe(":8080", nil)
}