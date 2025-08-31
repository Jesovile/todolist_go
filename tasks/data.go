package tasks

import (
	"github.com/google/uuid"
	"strconv"
)

// Tasks Repository Types

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type NewTaskData struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type TaskRepositoryI interface {
	GetAllTasks() []Task
	AddNewTask(newTask Task)
}

type taskRepositoryT struct {
	tasks []Task
}

// todo: delete this fake, for test only
func getStubData(count int) []Task {
	var result []Task
	for i := 1; i <= count; i++ {
		var stringIndex = strconv.Itoa(i)
		taskItem := Task{uuid.NewString(), "Task title " + stringIndex, "Task description for Task " + stringIndex, "todo"}
		result = append(result, taskItem)
	}
	return result
}

// GetAllTasks Tasks Repository Implementation
func (repo *taskRepositoryT) GetAllTasks() []Task {
	return repo.tasks
}
func (repo *taskRepositoryT) AddNewTask(newTask Task) {
	repo.tasks = append(repo.tasks, newTask)
}

// singleton instance
var repoInstance = &taskRepositoryT{
	tasks: getStubData(5),
}

func createRepositoryInstance() TaskRepositoryI {
	var repoInterface TaskRepositoryI = repoInstance
	return repoInterface
}

// TaskRepository - singleton pointer export as TaskRepositoryI interface instance
var TaskRepository = createRepositoryInstance()
