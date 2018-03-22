/**
 * INTRO
 * Jobs of a basic http server
 * 	1. Process dynamic requests.
 * 	2. Serve static assets
 * 	3. Accept connections
 *
 * PROCESS DYNAMIC REQUESTS
 * 	we can register a new handler with the http.HandleFunc function.
 * 	It's first parameter takes a path to match and a function to execute as a second
 *

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

 * 	The dynamic aspect comes from the http.Request, it contains all information about the request and it's parameters.
 * 	Read GET parameters with r.URL.Query().Get("token")
 * 	Read POST parameters with r.FormValue("email")
 *
 * SERVING STATIC ASSETS
 *  To serve static assets like JS, CSS, and images.
 *  Use the built in http.FileServer and point it to a url path.

    fs := http.FileServer(http.Dir("static/"))

 * Once our file server is in place, point a url path at it, just like we did with the dynamic requests.

    http.Handle("/static/", http.StripPrefix("/static/", fs))

 * ACCEPT CONNECTIONS
 *  Lastly, the HTTP server needs to listen on a port.
 *  Go has a built in HTTP server, we can start.

    http.ListenAndServe(":80", nil)

 *  pulling all together
 */
package main

import (
	"fmt"
	"io"
	"net/http"
)

func log(a *http.Request) {
	fmt.Println(a)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log(r)
	io.WriteString(w, "Hello")
}

func index(w http.ResponseWriter, r *http.Request) {
	log(r)
	fmt.Fprintf(w, "Benvenuto, alle umili origini.")
}

func main() {
	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}
