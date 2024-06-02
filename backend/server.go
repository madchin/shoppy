package main

import (
	"backend/data"
	db "backend/db"
	handler "backend/handlers"
	adapters "backend/internal/users/adapters/repository"
	"backend/internal/users/app"
	ports "backend/internal/users/port/http"
	"backend/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
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
	userRepo := data.NewUserRepository(DB)
	userHandler := handler.NewUser(userRepo).Build()

	logger := logrus.New()

	db2, _ := adapters.NewDatabase()
	app := app.NewApplication(adapters.NewUserRepository(db2), logrus.NewEntry(logger))
	httpServer := ports.NewHttpServer(app)
	http.HandleFunc("/api/v3/user", httpServer.User.Get)
	//http.Handle("/api/user", handler.SessionMiddleware(data.NewUserRepository(DB), handler.User))
	//http.Handle("/api/v2/user", middleware.AuthMiddleware(handler.User(data.NewUserRepository(DB))))
	http.Handle("/api/v2/user", middleware.AuthMiddleware(userHandler))
	// mux.HandleFunc("/api/v2/user", middleware.LoggerMiddleware(middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request, uuid string) {
	// 	handler.User(data.NewUserRepository(DB), uuid, w, r)
	// })))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil))
}
