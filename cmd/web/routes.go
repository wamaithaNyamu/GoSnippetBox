package main

import (
    "net/http"
	"github.com/julienschmidt/httprouter"

    "github.com/justinas/alice" // New import
)

// Update the signature for the routes() method so that it returns a
// http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {
   
	// Initialize the router.
    router := httprouter.New()
	   // Create a handler function which wraps our notFound() helper, and then
    // assign it as the custom handler for 404 Not Found responses. You can also
    // set a custom handler for 405 Method Not Allowed responses by setting
    // router.MethodNotAllowed in the same way too.
    router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.notFound(w)
    })



    // Update the pattern for the route for the static files.
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	   // Create a new middleware chain containing the middleware specific to our
    // dynamic application routes. For now, this chain will only contain the
    // LoadAndSave session middleware but we'll add more to it later.
    dynamic := alice.New(app.sessionManager.LoadAndSave)


     // Update these routes to use the new dynamic middleware chain followed by
    // the appropriate handler function. Note that because the alice ThenFunc()
    // method returns a http.Handler (rather than a http.HandlerFunc) we also
    // need to switch to registering the route using the router.Handler() method.
    router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
    router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
    router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
    router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))
	// Add the five new routes, all of which use our 'dynamic' middleware chain.
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))
	router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.userLogoutPost))


	//  Create a middleware chain containing our 'standard' middleware
    // which will be used for every request our application receives.
    standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    // Return the 'standard' middleware chain followed by the servemux.
    return standard.Then(router)
}