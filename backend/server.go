package main

import (
	adapters "backend/internal/users/adapters/repository"
	"backend/internal/users/app"
	ports "backend/internal/users/port/http"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	database, err := adapters.NewDatabase()
	if err != nil {
		panic(fmt.Sprintf("Database Initialization error", err.Error()))
	}
	defer database.Close()
	userRepository := adapters.NewUserRepository(database)
	userDetailRepository := adapters.NewUserDetailRepository(database)
	phoneRepository := adapters.NewPhoneRepository(database)
	addressRepository := adapters.NewAddressRepository(database)

	logger := logrus.New()
	app := app.NewApplication(userRepository, userDetailRepository, phoneRepository, addressRepository, logrus.NewEntry(logger))

	httpServer := ports.NewHttpServer(app)
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		panic("Environment variable \"httpPort\" not specified")
	}

	server := &http.Server{
		Handler: ports.HandlerFromMux(httpServer, http.NewServeMux()),
		Addr:    fmt.Sprintf(":%s", httpPort),
	}

	logger.Fatal(server.ListenAndServe())

}
