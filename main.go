package main

import (
	"log"
	"os"

	"github.com/hidayatullahap/go-todo-example/action"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", action.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("app_port")))
}
