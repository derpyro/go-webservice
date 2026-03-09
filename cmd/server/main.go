package main

import (
	"go-webservice/animal-service/internal/animal"
	"net/http"

	"fmt"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting Server...")
	store := animal.NewStore()
	service := animal.NewService(store)
	handler := animal.NewHandler(service)

	router := chi.NewRouter()
	handler.Routes(router)
	http.ListenAndServe(":8080", router)
}
