package servesTask

import "gorm.io/gorm"

type TaskNew struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
