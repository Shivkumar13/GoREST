package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// The "HandleFunc" method accepts a path and a function as arguments
	// Yes we can pass functions as arguments, and even treat them like variables in Go
	// However, the handler function has to have the appropriate signature (as desribed by the "handler" function below)
	router := httprouter.New()
	//router.Get("/", Index)
	router.GET("/hello/:name", Hello)

	//	http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 8080
	// The Second argument is the handler, which we will comet to later on, but now it is left as nil
	// and the handler defined above (in "HandleeFunc") is used
	http.ListenAndServe(":8080", router)

}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {

		//For this case, we will always pipe "Hello from Shivkumar Ople" into the response writer
		fmt.Fprint(w, "<h1>Hello From Shivkumar Ople!</h1>")

	} else if r.URL.Path == "/contact" {

		fmt.Fprintf(w, "Want to connect? please send an email to <a href=\"mailto:shivkumarople@gmail.com\">shiv@dracutt.com</a>.")
	} else {

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep 
	being sent to an invalid page.</p>`)
	}

	//	fmt.Fprint(w, r.URL.Path)

}
