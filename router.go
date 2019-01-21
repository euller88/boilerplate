package boilerplate

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter cria um novo roteador e aplica o middleware de log para nas rotas avaliáveis, além de implementar dois handlers padrões para erros 404 e 405. Precisa chamado depois de RegisterRoute(r *Route)
//
//NewRouter, well, creates a new router, and applies the logging middleware on the available routes, it also implements the default handler for the 404 and 405 errors. Must be called after RegisterRoute(r *Route)
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.Handler
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		log.Printf(
			"Route %s mounted, with the following path: %s %s\n",
			route.Name,
			route.Method,
			route.Pattern,
		)
	}

	router.NotFoundHandler = Logger(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)

			data := errorMessage{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Detail: "request not found",
			}

			json.NewEncoder(w).Encode(data)
		}),
		"DefaultNotFound",
	)

	router.MethodNotAllowedHandler = Logger(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusMethodNotAllowed)

			data := errorMessage{
				Code:   http.StatusMethodNotAllowed,
				Status: http.StatusText(http.StatusMethodNotAllowed),
				Detail: "method not allowed",
			}

			json.NewEncoder(w).Encode(data)
		}),
		"NotAllowed",
	)

	return router
}
