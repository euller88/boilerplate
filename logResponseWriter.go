package boilerplate

import "net/http"

//LoggingResponseWriter uma implementação de http.ResponseWriter capaz de expôr o seu código de status
//
//LoggingResponseWriter is an implementation of http.ResponseWriter able to expose its status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

//NewLoggingResponseWriter retorna um LoggingResponseWriter
//
//NewLoggingResponseWriter returns a LoggingResponseWriter
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

//WriteHeader é análogo a http.ResponseWriter.WriteHeader, porém expõe o código de status da requisição
//
//WriteHeader is analogous to http.ResponseWriter.WriteHeader, but exposes the status code of the request
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
