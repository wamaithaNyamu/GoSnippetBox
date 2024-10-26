package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// define a home handler (same as controller) which writes a byte slice
// containing " Hello from the other side" as the response body

func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello from the other side"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
    // Extract the value of the id parameter from the query string and try to
    // convert it to an integer using the strconv.Atoi() function. If it can't
    // be converted to an integer, or the value is less than 1, we return a 404 page
    // not found response.
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    // Use the fmt.Fprintf() function to interpolate the id value with our response
    // and write it to the http.ResponseWriter.
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed",  http.StatusMethodNotAllowed)
		return
	}
    w.Write([]byte("Create a new snippet..."))
}


func main(){
	// use the http.NewServeMux() function to initialize a new servemux then register
	// the handler above as the handler for the "/" url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)


    // Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen on (in this case ":4000")
    // and the servemux we just created. If http.ListenAndServe() returns an error
    // we use the log.Fatal() function to log the error message and exit. Note
    // that any error returned by http.ListenAndServe() is always non-nil.

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}