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
    boilerplate.OK("Hello World!",w)
}

func HelloVar(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    boilerplate.OK("Hello World!"+vars["v"] ,w)
}

func main() {
    boilerplate.RegisterRoute(&boilerplate.Route{
        Handler: http.HandlerFunc(Hello),
		Method:  "GET",
		Name:    "Hello",
		Pattern: "/",
    })
    
    boilerplate.RegisterRoute(&boilerplate.Route{
        Handler: http.HandlerFunc(HelloVar),
		Method:  "GET",
		Name:    "HelloVar",
		Pattern: "/{v}",
    })
    
    router := boilerplate.NewRouter()
    
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
    boilerplate.OK("Hello World!",w)
}

func HelloVar(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    boilerplate.OK("Hello World!"+vars["v"] ,w)
}

func init() {
    boilerplate.RegisterRoute(&boilerplate.Route{
        Handler: http.HandlerFunc(Hello),
		Method:  "GET",
		Name:    "Hello",
		Pattern: "/",
    })
    
    boilerplate.RegisterRoute(&boilerplate.Route{
        Handler: http.HandlerFunc(HelloVar),
		Method:  "GET",
		Name:    "HelloVar",
		Pattern: "/{v}",
    })
}
```
***
```go
package main

import (
    "log"
	"net/http"
	"github.com/euller88/boilerplate"
	//import like that in order to call the init() function of the package
	//importe assim para poder chamar a função init() do pacote
	_ "your/repository/name/handlers"
)

func main() {
    router := boilerplate.NewRouter()
    
    log.Fatalln(http.ListenAndServe(":8080", router))
}
```