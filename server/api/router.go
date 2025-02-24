package api

import (
	"github.com/PainCodermax/to-do-list-api/pkg/setting"
	v1 "github.com/PainCodermax/to-do-list-api/server/api/v1"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration) {
	task := v1.NewTask(cfg)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/task", task.Create)
		apiv1.DELETE("/task/:id", task.DeleteByID)
		apiv1.PUT("/task/:id", task.CompletedTaskByID)
		apiv1.GET("/tasks", task.GetAll)
		apiv1.GET("/task/:id", task.GetByID)
	}
}
