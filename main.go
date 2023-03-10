package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pedrosouza458/crud-go/configs"
	"github.com/pedrosouza458/crud-go/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	route := chi.NewRouter()

	route.Post("/", handlers.Create)
	route.Put("/{id}", handlers.Update)
	route.Delete("/{id}", handlers.Delete)
	route.Get("/", handlers.List)
	route.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), route)

}