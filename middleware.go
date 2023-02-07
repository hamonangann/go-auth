package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	USERNAME := os.Getenv("GOAUTH_USERNAME")
	PASSWORD := os.Getenv("GOAUTH_PASSWORD")
	requestUsername, requestPassword, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something were wrong`))
	}

	if isValid := (requestUsername == USERNAME) && (requestPassword == PASSWORD); !isValid {
		w.Write([]byte(`Wrong credentials!`))
		return false
	}
	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`Only GET is allowed`))
		return false
	}
	return true
}
