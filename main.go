package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomMux struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middleware = append(c.middleware, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	current := http.Handler(&c.ServeMux)
	for _, next := range c.middleware {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

func ConvertObjectToJSON(w http.ResponseWriter, o any) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func ActionUser(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		ConvertObjectToJSON(w, SelectUser(id))
		return
	}

	ConvertObjectToJSON(w, GetUsers())
}

func main() {
	mux := new(CustomMux)
	mux.HandleFunc("/user", ActionUser)

	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":9000" // listen to port 9000
	server.Handler = mux

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}
