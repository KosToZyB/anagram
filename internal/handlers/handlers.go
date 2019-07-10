package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kostozyb/anagram/internal/middleware"
	"github.com/kostozyb/anagram/pkg/anagram"
	"github.com/kostozyb/anagram/pkg/logger"
)

func NewAPI(log logger.Logger) *API {
	return &API{
		log: log,
		anagram: anagram.NewAnagram(),
	}
}

type API struct {
	log logger.Logger
	anagram *anagram.Anagram
}

func (a *API) Router() http.Handler {
	router := mux.NewRouter()

	subRouter := router.StrictSlash(true)

	mw := middleware.NewMiddleware(a.log)

	a.addRoutes(subRouter)
	subRouter.Use(mw.LoggerRequest)
	subRouter.Use(mw.LoggerResponse)

	return subRouter
}

func (a *API) addRoutes(r *mux.Router) {
	r.HandleFunc("/load", a.load).Methods(http.MethodPost)
	r.HandleFunc("/get", a.getAnagram).Methods(http.MethodGet)
}