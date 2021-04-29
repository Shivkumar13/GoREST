package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// The "HandleFunc" method accepts a path and a function as arguments
	// Yes we can pass functions as arguments, and even treat them like variables in Go
	// However, the handler function has to have the appropriate signature (as desribed by the "handler" function below)
	//router := httprouter.New()

	r := mux.NewRouter()
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/hello", Hello)
	r.HandleFunc("/contact", contact)

	http.NotFoundHandler()

	http.ListenAndServe(":8080", r)

	//router.Get("/", Index)

	// router.GET("/hello/:name", Hello)

	//router.GET("/faq", faq)

	//http.HandleFunc("/faqsite", faq)

	//	http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 8080
	// The Second argument is the handler, which we will comet to later on, but now it is left as nil
	// and the handler defined above (in "HandleeFunc") is used

}

func faq(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Frequently Asked Questions</h1><p>Here are the list of questions that our users commonly ask.</p>")

}

func Hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello, shiva")
}

func contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Want to connect? please send an email to <a href=\"mailto:shivkumarople@gmail.com\">shiv@dracutt.com</a>.")

}

// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.

// func handler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "text/html")

// 	if r.URL.Path == "/" {

// 		//For this case, we will always pipe "Hello from Shivkumar Ople" into the response writer
// 		fmt.Fprint(w, "<h1>Hello From Shivkumar Ople!</h1>")

// 	} else if r.URL.Path == "/contact" {

// 		fmt.Fprintf(w, "Want to connect? please send an email to <a href=\"mailto:shivkumarople@gmail.com\">shiv@dracutt.com</a>.")
// 	} else {

// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, `<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep
// 	being sent to an invalid page.</p>`)
// 	}

//	fmt.Fprint(w, r.URL.Path)
