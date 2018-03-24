/**
 * FORMS
 *   This example will show how to simulate a contact form an parse the message into a struct
 */
package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email: r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// Print details:
		fmt.Println(details)

		tmpl.Execute(w, struct{Success bool}{true})
	})

	http.ListenAndServe(":8080", nil)
}