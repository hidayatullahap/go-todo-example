package main

import (
	"log"
	"os"
	"strconv"

	"github.com/hidayatullahap/go-todo-example/action"
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
	e.Validator = core.NewCustomValidator()
}

func addRoutes(e *echo.Echo, app core.App) {
	todoAction := action.NewTodo(&app)
	tagAction := action.NewTag(&app)

	e.GET("/", action.Hello)

	e.GET("/todos", todoAction.FindList)
	e.GET("/todos/:id", todoAction.FindDetail)
	e.POST("/todos", todoAction.Create)
	e.PUT("/todos/:id", todoAction.Update)
	e.PATCH("/todos/:id", todoAction.UpdateStatus)

	e.GET("/tags", tagAction.FindList)
	e.GET("/tags/:id", tagAction.FindDetail)
	e.POST("/tags", tagAction.Create)
	e.PUT("/tags/:id", tagAction.Update)
}

func initApp() core.App {
	mysqlCon := database.GetMysqlConnection()
	debugIsActive, _ := strconv.ParseBool(os.Getenv("db_debug"))
	if debugIsActive {
		mysqlCon.LogMode(true)
	}

	return core.App{
		Db: mysqlCon,
	}
}
