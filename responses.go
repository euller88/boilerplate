package boilerplate

import (
	"encoding/json"
	"net/http"
)

//errorMessage representa um erro que pode ocorrer na aplicação
//
//errorMessage represents an error that might occur in the application
type errorMessage struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}

//UnprocessableEntity retorna uma resposta para o status HTTP 422
//
//UnprocessableEntity returnas a response to the HTTP status 422
func UnprocessableEntity(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnprocessableEntity)

	data := errorMessage{
		Code:   http.StatusUnprocessableEntity,
		Status: http.StatusText(http.StatusUnprocessableEntity),
		Detail: err.Error(),
	}

	json.NewEncoder(w).Encode(data)
}

//BadRequest retorna uma resposta para o status HTTP 400
//
//BadRequest returnas a response to the HTTP status 400
func BadRequest(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	data := errorMessage{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Detail: err.Error(),
	}

	json.NewEncoder(w).Encode(data)
}

//InternalServerError retorna uma resposta para o status HTTP 500
//
//InternalServerError returnas a response to the HTTP status 500
func InternalServerError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)

	data := errorMessage{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Detail: err.Error(),
	}

	json.NewEncoder(w).Encode(data)
}

//NotFound retorna uma resposta para o status HTTP 404
//
//NotFound returnas a response to the HTTP status 404
func NotFound(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	data := errorMessage{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Detail: err.Error(),
	}

	json.NewEncoder(w).Encode(data)
}

//Forbidden retorna uma resposta para o status HTTP 403
//
//Forbidden returnas a response to the HTTP status 403
func Forbidden(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)

	data := errorMessage{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
		Detail: err.Error(),
	}

	json.NewEncoder(w).Encode(data)
}

//OK retorna uma resposta para o status HTTP 200
//
//OK returnas a response to the HTTP status 200
func OK(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
