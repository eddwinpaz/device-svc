package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/eddwinpaz/device-svc/api"
	"github.com/eddwinpaz/device-svc/repository"
	"github.com/eddwinpaz/device-svc/service"
)

func main() {

	repo, _ := repository.NewPostgresRepository()
	service := service.NewLogService(repo)
	handler := api.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/ingest", handler.Post).Methods("POST")
	r.HandleFunc("/device/{id}", handler.Get).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	port := ":9000"
	fmt.Printf("Server running on port %s\n", port)

	err := http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(r))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

}
