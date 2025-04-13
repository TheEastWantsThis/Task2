package servesTask

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(tas TaskNew) error
	GetAllTasks() ([]TaskNew, error)
	GetTaskById(id string) (TaskNew, error)
	UpdateTask(tas TaskNew) error
	DeleteTask(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(tas TaskNew) error {
	return r.db.Create(&tas).Error
}
func (r *taskRepository) GetAllTasks() ([]TaskNew, error) {
	var tasks []TaskNew
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) GetTaskById(id string) (TaskNew, error) {
	var task TaskNew
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}
func (r *taskRepository) UpdateTask(tas TaskNew) error {
	return r.db.Save(&tas).Error
}
func (r *taskRepository) DeleteTask(id string) error {
	return r.db.Delete(&TaskNew{}, "id = ?", id).Error
}
