/**
 * INTRO
 *   Go's html/templates package provides a rich templating lang for HTML templates.
 *
 * FIRST TEMPLATE:
 *   Data passed in, can be any kind of Go's data structures.
 *   To access the data in a template the top most variable is access by {{.}}.
 *   The dot inside the curly braces is called the pipeline and the root element of the data.
 *

	GO DATA:
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{ Title: "Task 1", Done: false},
			{ Title: "Task 2", Done: false},
			{ Title: "Task 3", Done: false},
		},
	}

	HTML TEMPLATE:
	<h1>{{.PageTitle}}</div>
	<ul>
		{{range .Todos}}
			{{if .Done}}
				<li class="done">{{.Title}}</li>
			{{else}}
				<li>{{.Title}}</li>
			{{end}}
		{{end}}
	</ul>

 * CONTROL STRUCTURES
 *   Short list of the most common control structures.

	 Control Structure	    			Definition
	 {{.}}	                			Renders the root element
	 {{.Title}}							Renders the “Title”-field in a nested element
	 {{if .Done}} {{else}} {{end}}		Defines an if-Statement
	 {{range .Todos}} {{.}} {{end}}		Loops over all “Todos” and renders each using {{.}}
	 {{block "content" .}} {{end}}		Defines a block with the name “content”

 * PARSING TEMPLATES FROM FILES
 *   Templates can be parsed from a string or a file on disk.

	 tmpl, err := template.ParseFiles("layout.html")
	 // or
	 tmpl := template.Must(template.ParseFiles("layout.html"))

 * EXECUTE A TEMPLATE IN A REQUEST HANDLER
 *   Once the template if parsed from disk it's ready to be used in the request handler.
 *   The 'Execute' function accepts an 'io.writer' for writing out the template and an 'interface{}' to pass data into the template.
 *   When the function is called on an 'http.ResponseWriter' the Content-Type in the header is automatically set in the HTTP response to 'Content-Type: text/html; charset=utf-8'

	func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, "data goes here")
 	}

 */
package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}