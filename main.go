package main

import (
	"log"
	"net/http"
)

// define a home handler (same as controller) which writes a byte slice
// containing " Hello from the other side" as the response body

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from the other side"))
}

func main(){
	// use the http.NewServeMux() function to initialize a new servemux then register
	// the handler above as the handler for the "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)

    // Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen on (in this case ":4000")
    // and the servemux we just created. If http.ListenAndServe() returns an error
    // we use the log.Fatal() function to log the error message and exit. Note
    // that any error returned by http.ListenAndServe() is always non-nil.

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}