package tasks

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
)

func getAllTasks(context *gin.Context) {
	tasks := TaskRepository.GetAllTasks()
	context.JSON(http.StatusOK, tasks)
}

func addNewTask(context *gin.Context) {
	var newTaskData = NewTaskData{
		Title:       "",
		Description: "",
	}
	if err := context.ShouldBindBodyWith(&newTaskData, binding.JSON); err != nil {
		fmt.Println("Error during adding new task", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newTask = Task{
		Id:          uuid.NewString(),
		Title:       newTaskData.Title,
		Description: newTaskData.Description,
		Status:      "todo",
	}
	TaskRepository.AddNewTask(newTask)
	context.String(http.StatusOK, "Task is added")
}

func SetTasksApi(router *gin.Engine) {
	router.GET("/task", getAllTasks)
	router.POST("task", addNewTask)
}
