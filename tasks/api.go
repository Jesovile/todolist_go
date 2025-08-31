package tasks

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllTasks(context *gin.Context) {
	tasks := TaskRepository.GetAllTasks()
	context.JSON(http.StatusOK, tasks)
}

func SetTasksApi(router *gin.Engine) {
	router.GET("/tasks", getAllTasks)
}
