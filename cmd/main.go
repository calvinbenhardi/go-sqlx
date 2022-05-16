package main

import (
	"fmt"
	"net/http"

	"github.com/calvinbenhardi/go-sqlx/internal/delivery/rest"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", rest.HealthCheck).Methods(http.MethodGet)

	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Print(err)
	}
}
