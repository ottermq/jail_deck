package main

import (
	"log"
	"net/http"

	"github.com/otterlabs/jaildeck/internal/app"
)

func main() {
	a := app.New()
	log.Printf("Jail Deck listening on http://127.0.0.1:8888")
	err := http.ListenAndServe(":8888", a.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
