package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sanchezqb/go-rest-api/todo"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //Set Content-Type Headers as application/json
		middleware.Logger,          //Log API requests
		middleware.RedirectSlashes, //Redirect trailing slashes to no trailing slash url
		middleware.Recoverer,       //Recover from panics without crashing
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.Routes())
	})

	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) //Walk and print all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error()) //Panic if there was an error
	}

	log.Fatal(http.ListenAndServe(":8000", router))
}
