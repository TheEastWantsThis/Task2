package servesTask

type TaskNew struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func (TaskNew) TableName() string {
	return "tasks"
}
