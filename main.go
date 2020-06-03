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
	addMiddleware(e)
	addRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("app_port")))
}

func addMiddleware(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func addRoutes(e *echo.Echo) {
	todoAction := action.NewTodo()

	e.GET("/", action.Hello)
	e.GET("/todos", todoAction.FindTodoList)
}
