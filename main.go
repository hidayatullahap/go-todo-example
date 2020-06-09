package main

import (
	"log"
	"os"

	"github.com/hidayatullahap/go-todo-example/action"
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := initApp()

	e := echo.New()
	addMiddleware(e)
	addRoutes(e, app)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("app_port")))
}

func addMiddleware(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func addRoutes(e *echo.Echo, app core.App) {
	e.GET("/", action.Hello)
}

func initApp() core.App {
	return core.App{}
}
