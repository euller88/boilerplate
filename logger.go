package boilerplate

import (
	"log"
	"net/http"
	"time"
)

//Logger é um middleware para "imprimir" as mensagens de log da aplicação
//
//Logger is a middleware to "print" the application's log messages
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lwr := NewLoggingResponseWriter(w)

		inner.ServeHTTP(lwr, r)

		log.Printf(
			"%s %s %d %s %s\n",
			name,
			r.Method,
			lwr.StatusCode,
			r.RequestURI,
			time.Since(start),
		)
	})
}
