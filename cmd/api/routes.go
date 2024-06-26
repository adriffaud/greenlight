package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.GET("/v1/healthcheck", app.healthcheckHandler)
	router.GET("/v1/movies", app.listMoviesHandler)
	router.POST("/v1/movies", app.createMovieHandler)
	router.GET("/v1/movies/:id", app.showMovieHandler)
	router.PATCH("/v1/movies/:id", app.updateMovieHandler)
	router.DELETE("/v1/movies/:id", app.deleteMovieHandler)
	router.POST("/v1/users", app.registerUserHandler)
	router.PUT("/v1/users/activated", app.activateUserHandler)

	return app.recoverPanic(app.rateLimit(router))
}
