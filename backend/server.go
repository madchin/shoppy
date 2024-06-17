package main

import (
	"backend/internal/common/auth"
	"backend/internal/common/server"
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
		panic(fmt.Sprintf("Database Initialization error %s", err.Error()))
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
		panic("Environment variable \"HTTP_PORT\" not specified")
	}
	privateKeyPath := os.Getenv("PRIVATE_USERS_API_KEY_PATH")
	if privateKeyPath == "" {
		panic("Environment variable \"PRIVATE_USERS_API_KEY_PATH\" not specified")
	}
	publicKeyPath := os.Getenv("PUBLIC_USERS_API_KEY_PATH")
	if publicKeyPath == "" {
		panic("Environment variable \"PUBLIC_USERS_API_KEY_PATH\" not specified")
	}
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read private key from PEM file, err: %s", err.Error()))
	}

	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read public key from PEM file, err: %s", err.Error()))
	}

	jwtAuth := auth.NewJwtAuth(privateKey, publicKey)
	httpHandler := ports.HandlerWithOptions(httpServer, ports.StdHTTPServerOptions{
		BaseURL:     "/api/v1",
		Middlewares: []ports.MiddlewareFunc{ports.MiddlewareFunc(server.AuthMiddleware(jwtAuth))},
	})
	server := &http.Server{
		Handler: httpHandler,
		Addr:    fmt.Sprintf(":%s", httpPort),
	}

	logger.Fatal(server.ListenAndServe())

}
