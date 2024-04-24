package main

import (
	data "backend/data"
	db "backend/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	DB, err := db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer DB.Close()

	e.GET("/", func(c echo.Context) error {
		return rootHandler(c, DB)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func rootHandler(c echo.Context, db *sql.DB) error {
	user := &data.User{Name: "elo", Email: "yoyo"}
	err := user.Create(db)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("Hello, Docker! (%v)\n", user))
}
