package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		ConvertObjectToJSON(w, SelectUser(id))
		return
	}

	ConvertObjectToJSON(w, GetUsers())
}

func main() {
	http.HandleFunc("/user", ActionUser)

	server := new(http.Server)
	server.Addr = ":9000" // listen to port 9000

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}