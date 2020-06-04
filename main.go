package main

import (
	"log"
	"os"

	"github.com/hidayatullahap/go-todo-example/action"
	"github.com/hidayatullahap/go-todo-example/action/request"
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := getEnv()

	e := echo.New()
	addMiddleware(e)
	addRoutes(e, env)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("app_port")))
}

func addMiddleware(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = request.NewCustomValidator()
}

func addRoutes(e *echo.Echo, env core.Env) {
	todoAction := action.NewTodo(&env)
	tagAction := action.NewTag(&env)

	e.GET("/", action.Hello)

	e.GET("/todos", todoAction.FindList)
	e.GET("/todos/:id", todoAction.FindDetail)
	e.POST("/todos", todoAction.Create)
	e.PUT("/todos/:id", todoAction.Update)

	e.GET("/tags", tagAction.FindList)
	e.GET("/tags/:id", tagAction.FindDetail)
	e.POST("/tags", tagAction.Create)
	e.PUT("/tags/:id", tagAction.Update)
}

func getEnv() core.Env {
	mysqlCon := database.GetMysqlConnection()

	return core.Env{
		Db: mysqlCon,
	}
}
