//A web application in Go.
//Author: Danielis Joniskis

package main

import (
//   "fmt"
	"net/http"
    "text/template"
    "strconv"
    "time"
    "math/rand"
)

//A http.ResponseWriter assembles the HTTP server's response by writing to it.
//A http.Request is a data structure that represents the client HTTP request.
/*func handler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type","text/html")
    fmt.Fprintln(w, "<h1>Guessing game</h1>")
}*/

//A struct
type Template struct {
    Message string
    Count int 
    Prompt string
    Guess int
}

//Adapted from https://www.youtube.com/watch?v=GTSq1VPPFco&feature=youtu.be.
func templateHandler(w http.ResponseWriter, r *http.Request){

        //Giving a random seed so that the same value is not generated every time.
	    rand.Seed(time.Now().UnixNano())
        //Generating a random int just like in part 5 of the first problem sheet.
        i := rand.Intn(20)
    
        //Reading the cookie //Adapted from https://github.com/data-representation/go-cookies
        var target, err = r.Cookie("i")
        if err == nil {
            // converting the value to an int
            i, _ = strconv.Atoi(target.Value)
        }

    //Setting an expiration date of a year
    expiration := time.Now().AddDate(1, 0, 0)
    

    //Creating and setting a session cookie
	target = &http.Cookie{
        Name:    "i",
        Expires: expiration,
		Value:   strconv.Itoa(i),
	}
    http.SetCookie(w, target)
    

    guess,_ := strconv.Atoi(r.FormValue("guess"))
    prompt :=""
    //An if else statement which compares the value of the guess to random number i 
    if guess == i{
        prompt ="Correct!"
    }else if guess < i{
        prompt="Guess higher"
     }else {
        prompt="Guess lower"
    }


    t, _ := template.ParseFiles("template/guess.html")//parsing the file and then executing the templates
    t.Execute(w, Template{Message: "Guess a number between 1 and 20",Prompt:prompt})
}

//The main function begins with a call to http.Handle.
//It then calls http.ListenAndServe, specifying that it should listen on port 8080.
func main() {
    http.Handle("/",  http.FileServer(http.Dir("./static")))//serving index.html in the root directory.
    http.HandleFunc("/guess", templateHandler)
    http.ListenAndServe(":8080", nil)
}