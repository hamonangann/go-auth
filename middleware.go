package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err.Error())
		}

		USERNAME := os.Getenv("GOAUTH_USERNAME")
		PASSWORD := os.Getenv("GOAUTH_PASSWORD")
		requestUsername, requestPassword, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something were wrong`))
			return
		}

		if isValid := (requestUsername == USERNAME) && (requestPassword == PASSWORD); !isValid {
			w.Write([]byte(`Wrong credentials!`))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte(`Only GET is allowed`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
