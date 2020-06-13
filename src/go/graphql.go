package main

import (
	"github.com/graphql-go/handler"
	"github.com/michailsmith/koan/data"
	"log"
	"net/http"
)

func main() {
	h := handler.New(&handler.Config{
		Pretty: true,
		Schema: &schema.Schema,
	})

	http.Handle("/graphql", h)

	port := ":8080"
	log.Printf("listening on port %v", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalf("listening failed, %v", err)
	}
}
