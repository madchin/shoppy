package main

import (
	db "backend/db"
	handler "backend/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	DB, err := db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer DB.Close()

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	http.Handle("/api/user", handler.SessionMiddleware(DB, handler.User))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil))
}
