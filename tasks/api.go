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
		fmt.Println("Error during deleting task "+id, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, "Task is deleted")
}

func updateTask(context *gin.Context) {
	id := context.Param("id")
	var updatedTask Task
	if err := context.ShouldBindBodyWith(&updatedTask, binding.JSON); err != nil {
		fmt.Println("Error during binding updated task "+id, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err2 := TaskRepository.UpdateTaskById(updatedTask); err2 != nil {
		fmt.Println("Error during updating task "+id, err2)
		context.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	context.String(http.StatusOK, "Task is updated")
}

func SetTasksApi(router *gin.Engine) {
	router.GET("/task", getAllTasks)
	router.POST("task", addNewTask)
	router.DELETE("/task/:id", deleteTaskById)
	router.PATCH("/task/:id", updateTask)
}
