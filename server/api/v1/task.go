package v1

import (
	"github.com/PainCodermax/to-do-list-api/internal/service"
	"github.com/PainCodermax/to-do-list-api/internal/service/serializers"
	"github.com/PainCodermax/to-do-list-api/pkg/app"
	"github.com/PainCodermax/to-do-list-api/pkg/errcode"
	"github.com/PainCodermax/to-do-list-api/pkg/logger"
	"github.com/PainCodermax/to-do-list-api/pkg/setting"
	"github.com/gin-gonic/gin"
)

type Task struct {
	cfg *setting.Configuration
}

func NewTask(cfg *setting.Configuration) Task {
	return Task{
		cfg: cfg,
	}
}

func (t *Task) Create(c *gin.Context) {
	requestBody := serializers.TaskCreateRequest{}
	response := app.NewResponse(c)
	if app.Validation(c, &requestBody, response) != nil {
		return
	}

	task, err := service.NewTaskService(c.Request.Context()).Create(&requestBody)
	if err != nil {
		logger.WithTrace(c).Errorf("service.NewTaskService().Create err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTaskFail)
		return
	}

	response.ToResponse(serializers.TaskCreateResponse{ID: task.ID})
}

func (t *Task) GetAll(c *gin.Context) {
	response := app.NewResponse(c)
	tasks, err := service.NewTaskService(c.Request.Context()).GetAll()
	if err != nil {
		logger.WithTrace(c).Errorf("service.NewTaskService().GetAll err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTaskListFail)
		return
	}
	taskItems := make([]serializers.TaskGetAllResponseItem, len(tasks))
	for i, task := range tasks {
		taskItems[i] = serializers.TaskGetAllResponseItem{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		}
	}
	response.ToResponse(serializers.TaskGetAllResponse{Tasks: taskItems})
}

func (t *Task) GetByID(c *gin.Context) {
	response := app.NewResponse(c)
	id := c.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("id is required"))
		return
	}
	task, err := service.NewTaskService(c.Request.Context()).GetByID(id)
	if err != nil {
		logger.WithTrace(c).Errorf("service.NewTaskService().GetByID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTaskFail)
		return
	}
	response.ToResponse(serializers.TaskGetByIdResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	})
}

func (t *Task) CompletedTaskByID(c *gin.Context) {
	response := app.NewResponse(c)
	id := c.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("id is required"))
		return
	}
	taskService := service.NewTaskService(c.Request.Context())
	if _, err := taskService.GetByID(id); err != nil {
		logger.WithTrace(c).Errorf("service.NewTaskService().GetByID err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTaskFail)
		return
	}
	if err := taskService.Complete(id); err != nil {
		logger.WithTrace(c).Errorf("service.NewTaskService().UpdateTaskByID err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTaskFail)
		return
	}
	response.ToResponse(gin.H{"completed": true})
}

func (t *Task) DeleteByID(c *gin.Context) {
	response := app.NewResponse(c)
	id := c.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("id is required"))
		return
	}
	if err := service.NewTaskService(c.Request.Context()).Delete(id); err != nil {
		logger.WithTrace(c).Errorf("service.NewTaskService().Delete err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTaskFail)
		return
	}
	response.ToResponse(gin.H{"deleted": true})
}
