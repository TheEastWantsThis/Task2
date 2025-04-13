package servesTask

type TaskServes interface {
	CreateTask(task TaskNew) (TaskNew, error)
	GetAllTasks() ([]TaskNew, error)
	GetTaskById(id string) (TaskNew, error)
	UpdateTask(id, task string, IsDone bool) (TaskNew, error)
	DeleteTask(id string) error
}

type servesTask struct {
	repo TaskRepository
}

func NewServesTask(r TaskRepository) TaskServes {
	return &servesTask{repo: r}
}

func (s *servesTask) CreateTask(task TaskNew) (TaskNew, error) {
	t := TaskNew{
		Task:   task.Task,
		IsDone: task.IsDone,
	}
	if err := s.repo.CreateTask(t); err != nil {
		return t, err
	}
	return t, nil
}

func (s servesTask) GetAllTasks() ([]TaskNew, error) {
	return s.repo.GetAllTasks()
}

func (s *servesTask) GetTaskById(id string) (TaskNew, error) {
	return s.repo.GetTaskById(id)
}

func (s *servesTask) UpdateTask(id, task string, IsDone bool) (TaskNew, error) {
	t, err := s.repo.GetTaskById(id)
	if err != nil {
		return TaskNew{}, err
	}
	t.Task = task
	t.IsDone = IsDone

	if err := s.repo.UpdateTask(t); err != nil {
		return TaskNew{}, err
	}
	return t, nil
}

func (s *servesTask) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
