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
	type NewTaskData struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}
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

func deleteTaskById(context *gin.Context) {
	id := context.Param("id")
	if err := TaskRepository.DeleteTaskById(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, "Task is deleted")
}

func SetTasksApi(router *gin.Engine) {
	router.GET("/task", getAllTasks)
	router.POST("task", addNewTask)
	router.DELETE("/task/:id", deleteTaskById)
}
