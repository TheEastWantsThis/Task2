package main

import (
	"awesomeProject1/internal/Hadlers"
	"awesomeProject1/internal/db"
	"awesomeProject1/internal/servesTask"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Не смоги подключиться к БД: %v", err)
	}

	e := echo.New()

	taskRepo := servesTask.NewTaskRepository(database)
	taskServies := servesTask.NewServesTask(taskRepo)
	taskHandler := Hadlers.NewTaskHandler(taskServies)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/task", taskHandler.GetHandler)
	e.POST("/task", taskHandler.PostHandler)
	e.PATCH("/task/:id", taskHandler.PatchHandler)
	e.DELETE("/task/:id", taskHandler.DeleteHandler)
	e.Start("localhost:8080")
}
