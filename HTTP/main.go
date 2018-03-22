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
 *
 */
package HTTP
