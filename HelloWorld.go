/**
 * GO IS A BATTERY INCLUDED PROGRAMMING LANGUAGE AND HAS A WEBSERVER BUILT IN.
 */
package main

import (
	"net/http"
	"fmt"
)

/**
 * REGISTERING A REQUEST HANDLER:
 * First, create a Handler which receives all incomming HTTP connections from browsers, HTTP clients or API requests.
 * A handler in Go is a function with this signature:
 */
 //func (w http.ResponseWriter, r *http.Request)
/**
 * The func receives two parameters:
 * 1. An http.ResponseWriter: which is where you write your text/html response to.
 * 2. An http.Request: which contains all information about this HTTP request including things like URL or header fields.
 */
// simple request handleer to defualt HTTP server
//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//})
/**
 * LISTEN FOR HTTP CONNECTIONS
 * The request handler alone can't accept any HTTP connections from the outside.
 * An HTTP server has to listen on a port to pass connections on to the request handler.
 */
 //http.ListenAndServer(":80", nil)

 /**
  * PULLING IT ALL TOGETHER
  */
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}