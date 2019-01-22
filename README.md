# Boilerplate
###### English
Boilerplate is a package that contains the set of code I always write when I need to create a router for a small REST API that sends JSON as a response. Much of the work, such as handling the request body and URL variables, is done by the handlers I implement, but that's not the focus of the package. I really tried to be idiomatic as possible, and the only package outside the standard library is the excellent [gorilla/mux](https://github.com/gorilla/mux).
###### Português
Boilerplate é um pacote que contém o conjunto de código que eu sempre escrevo quando eu necessito criar um roteador para uma pequena API REST que envie JSON como resposta. Boa parte do trabalho, como tratamento do corpo da requisição e das variáveis de URL, é feito pelos handlers que eu implemento, mas isso não é o foco do pacote. Eu realmente tentei ser o mais idiomático possível, e o único pacote fora da biblioteca padrão é o excelente [gorilla/mux](https://github.com/gorilla/mux).
### How to use / Como usar
##### Handlers inside the main package / Handlers dentro do pacote principal
```go
package main

import (
  "log"
	"net/http"
	"github.com/euller88/boilerplate"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
  boilerplate.OK("Hello World!", w)
}

func HelloVar(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  boilerplate.OK("Hello World! "+vars["v"], w)
}

func HelloJSONMap(w http.ResponseWriter, r *http.Request) {
  m := map[string]interface{}{"message":"Hello World"}
  boilerplate.OK(m, w)
}

func HelloJSONStruct(w http.ResponseWriter, r *http.Request) {
  s := &struct{
    Message string `json:"message"`
  }{
    "Hello world",
  }
  boilerplate.OK(s, w)
}

func HelloEmpty(w http.ResponseWriter, r *http.Request){
  boilerplate.OK("", w)
}

func main() {
  routes := make([]*boilerplate.Route, 0)

  routes = append(
    routes, 
    &boilerplate.Route{
      Handler: http.HandlerFunc(Hello),
      Method:  "GET",
	  	Name:    "Hello",
      Pattern: "/",
    },
    &boilerplate.Route{
      Handler: http.HandlerFunc(HelloVar),
      Method:  "GET",
	  	Name:    "HelloVar",
      Pattern: "/{v}",
    },
    &boilerplate.Route{
      Handler: http.HandlerFunc(HelloJSONMap),
      Method:  "GET",
      Name:    "HelloJSONMap",
      Pattern: "/map",
    },
    &boilerplate.Route{
      Handler: http.HandlerFunc(HelloJSONStruct),
      Method:  "GET",
      Name:    "HelloJSONStruct",
      Pattern: "/struct",
    },
  )

  router := boilerplate.NewRouter()

  router.AddRoute(routes...)

  r := &boilerplate.Route{
    Handler: http.HandlerFunc(HelloEmpty),
    Method:  "GET",
    Name:    "HelloEmpty",
    Pattern: "/empty",
  },

  router.AddRoute(r)    
    
  log.Fatalln(http.ListenAndServe(":8080", router))
}
```
##### Handlers outside the main package / Handlers fora do pacote principal
```go
package handlers

import (
	"net/http"
	"github.com/euller88/boilerplate"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
  boilerplate.OK("Hello World!", w)
}

func HelloVar(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  boilerplate.OK("Hello World! "+vars["v"], w)
}

func HelloJSONMap(w http.ResponseWriter, r *http.Request) {
  m := map[string]interface{}{"message":"Hello World"}
  boilerplate.OK(m, w)
}

func HelloJSONStruct(w http.ResponseWriter, r *http.Request) {
  s := &struct{
    Message string `json:"message"`
  }{
    "Hello world",
  }
  boilerplate.OK(s, w)
}

func HelloEmpty(w http.ResponseWriter, r *http.Request){
  boilerplate.OK("", w)
}
```
***
```go
package main

import (
  "log"
	"net/http"
	"github.com/euller88/boilerplate"
	"your/repository/name/handlers"
)

func main() {
  routes := make([]*boilerplate.Route, 0)

  routes = append(
    routes, 
    &boilerplate.Route{
      Handler: http.HandlerFunc(handlers.Hello),
      Method:  "GET",
	  	Name:    "Hello",
      Pattern: "/",
    },
    &boilerplate.Route{
      Handler: http.HandlerFunc(handlers.HelloVar),
      Method:  "GET",
	  	Name:    "HelloVar",
      Pattern: "/{v}",
    },
    &boilerplate.Route{
      Handler: http.HandlerFunc(handlers.HelloJSONMap),
      Method:  "GET",
      Name:    "HelloJSONMap",
      Pattern: "/map",
    },
    &boilerplate.Route{
      Handler: http.HandlerFunc(handlers.HelloJSONStruct),
      Method:  "GET",
      Name:    "HelloJSONStruct",
      Pattern: "/struct",
    },
  )

  router := boilerplate.NewRouter()

  router.AddRoute(routes...)

  r := &boilerplate.Route{
    Handler: http.HandlerFunc(handlers.HelloEmpty),
    Method:  "GET",
    Name:    "HelloEmpty",
    Pattern: "/empty",
  },

  router.AddRoute(r)    
    
  log.Fatalln(http.ListenAndServe(":8080", router))
}
```