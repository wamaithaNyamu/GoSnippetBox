package main

import (
    "net/http"
	"github.com/julienschmidt/httprouter"

    "github.com/justinas/alice" // New import
)

func (app *application) routes() http.Handler {
    router := httprouter.New()

    router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        app.notFound(w)
    })
    
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// Add the authenticate() middleware to the chain.
    dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

    router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
    router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
    router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
    router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
    router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
    router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

    // Because the 'protected' middleware chain appends to the 'dynamic' chain
    // the noSurf middleware will also be used on the three routes below too.
    protected := dynamic.Append(app.requireAuthentication)

    router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
    router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
    router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

    standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
    return standard.Then(router)
}