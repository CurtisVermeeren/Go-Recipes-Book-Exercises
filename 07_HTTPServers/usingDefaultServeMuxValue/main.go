package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html>
 <html>
 <head>
 <title>Hello Gopher</title> 
 </head>
 <body>
 <b>Hello Gopher!</b>
 <p>
 <a href="/welcome">Welcome</a> | <a href="/message">Message</a>
 </p>
 </body>
</html>`
	fmt.Fprintf(w, html)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Programming")
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "net/http package is used to build web apps")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/message", message)
	log.Println("Listening...")
	// using nil will use the DefaultServeMUX
	// you can use the function http.HandleFunc to register
	// a handler function for the given URL pattern. Inside the function http.HandleFunc , it calls the function
	// HandleFunc of DefaultServeMux a handler function for the given URL pattern
	http.ListenAndServe(":8080", nil)
}
