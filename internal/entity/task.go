package entity

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var TaskList []Task

func GetTaskList() []Task {
	return TaskList
}
