package main

import (
	"log"
	"net/http"

	"github.com/ajorgensen/go-svelte-test/handler"
)

func main() {
	router := handler.NewRouter()

	router.ServeFiles("/static/*filepath", http.Dir("public"))

	router.Get("/", handler.Handler{H: handler.Index})
	router.Get("/blog", handler.Handler{H: handler.Blog})

	log.Fatal(http.ListenAndServe(":8000", router))
}
