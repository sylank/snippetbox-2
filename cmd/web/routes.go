package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddlewares := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddlewares.ThenFunc(http.HandlerFunc(app.home)))
	mux.Get("/snippet/create", dynamicMiddlewares.ThenFunc(http.HandlerFunc(app.createSnippetForm)))
	mux.Post("/snippet/create", dynamicMiddlewares.ThenFunc(http.HandlerFunc(app.createSnippet)))
	mux.Get("/snippet/:id", dynamicMiddlewares.ThenFunc(http.HandlerFunc(app.showSnippet)))

	mux.Get("/user/signup", dynamicMiddlewares.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddlewares.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddlewares.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddlewares.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddlewares.ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
