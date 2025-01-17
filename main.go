package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// define a home hanlder function which writes a byte slice contiaining
// "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	//extract the value of the id wildcard from the request using r.PathValue()
	//and try ton convert it to an integer using the strconv.Atoi() function. if
	//it can't be converted to an integer, or the value is less than 1, we
	//return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// use fmt.Sprintf() function to interpolate the id value with a
	// message, then write it as the HTTP response.
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))

}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form creating  new snippet..."))
}

func main() {

	//Use the http.NewServeMux() funcion to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView) //add the {id} wildcard segment
	mux.HandleFunc("/snippet/create", snippetCreate)

	//print a log a message to say that the server is starting.
	log.Print("starting server on :4000")

	//use the http.ListenAndServer() function to start a new web server. we pass in
	//two parameters: the TCP network address to listen on (in this case ":4000")
	//and the servemux we just created. if http.ListenAndServe() returns an error
	//that any error returned by http.ListenAndServe() is always non-nill.

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
