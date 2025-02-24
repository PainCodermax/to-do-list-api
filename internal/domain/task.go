package domain

import (
	"errors"
	"fmt"

	"github.com/PainCodermax/to-do-list-api/internal/entity"
)

type TaskDomain struct {
	Tasks []entity.Task
}

func NewTaskDomain() *TaskDomain {
	return &TaskDomain{Tasks: entity.GetTaskList()}
}

func (d *TaskDomain) Create(id, desscription, title string) (*entity.Task, error) {
	task := entity.Task{
		ID:          id,
		Description: desscription,
		Title:       title,
		Completed:   false,
	}
	entity.TaskList = append([]entity.Task{task}, d.Tasks...)
	return &task, nil
}

func (d *TaskDomain) GetAll() []entity.Task {
	return d.Tasks
}

func (d *TaskDomain) GetByID(id string) (*entity.Task, error) {
	for _, task := range d.Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (d *TaskDomain) Complete(id string) (*entity.Task, error) {
	for i, task := range d.Tasks {
		if task.ID == id {
			d.Tasks[i].Completed = true
			return &d.Tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (d *TaskDomain) Delete(id string) error {
	for i, task := range d.Tasks {
		if task.ID == id {
			entity.TaskList = append(d.Tasks[:i], d.Tasks[i+1:]...)
			return nil
		}
	}
	fmt.Println(d.Tasks)
	return errors.New("task not found")
}
