package Hadlers

import (
	"awesomeProject1/internal/servesTask"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	serves servesTask.TaskServes
}

func NewTaskHandler(s servesTask.TaskServes) *TaskHandler {
	return &TaskHandler{serves: s}
}

func (h *TaskHandler) GetHandler(c echo.Context) error {
	task, err := h.serves.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не смоги получить задания"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var task servesTask.TaskNew
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Hекорректные данные"})
	}

	t, err := h.serves.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Hе смоги создать задание"})
	}

	return c.JSON(http.StatusCreated, t)

}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	id := c.Param("id")

	var req servesTask.TaskNew
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Не то написал"})
	}

	updateTask, err := h.serves.UpdateTask(id, req.Task, req.IsDone)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Не можем обновить вычисления"})
	}

	return c.JSON(http.StatusOK, updateTask)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if err := h.serves.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не смоги удалить задание"})
	}

	return c.NoContent(http.StatusNoContent)
}
