/**
 * INTRO
 * Gorilla/mux is the community router of choice for adding named parameters.
 *
 * INSTALL
 	go get -u github.com/gorilla/mux

 * CREATING A NEW ROUTER
 *   Create a new router

	 r := mux.NewRouter()

 * REGISTERING A REQUEST HANDLER
 *   Once a router has been created.
 *   Register request with r.HandleFunc, instead of http.HandleFunc

	 r.HandleFunc

 * URL PARAMETERS
 *  Extract segments out of the url
 *  i.e: /books/go-programming-blueprint/page/10 => /books/{title}/page/{page}
 *
 *  To have a request handler match the URL mentioned above you replace the dynamic segments of the URL with placeholders.
 *  Then to get the vars use mux.Vars(r) that returns a map of the segments.

	f.HandleFunc("/books/{title}/pages/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vars["title"] // => the book title slug
		vars["page"]  // => the page slug
	}

 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars  := mux.Vars(r)
		title := vars["title"]
		page  := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
		fmt.Println("VARS:", title, page)
	})

	http.ListenAndServe(":8000", r)
}

