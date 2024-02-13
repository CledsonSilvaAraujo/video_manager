package service

import (
	"github.com/CledsonSilvaAraujo/task-manager/entity"
)

type TaskService interface {
	Save(entity.Task) entity.Task
	FindAll() []entity.Task
}

type taskService struct {
	tasks []entity.Task
}

func NewTaskService() TaskService {
	return &taskService{}
}

func (service *taskService) Save(task entity.Task) entity.Task {
	service.tasks = append(service.tasks, task)
	return task
}

func (service *taskService) FindAll() []entity.Task {
	return service.tasks
}
