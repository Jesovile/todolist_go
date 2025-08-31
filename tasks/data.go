package tasks

import (
	"errors"
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

type TaskRepositoryI interface {
	GetAllTasks() []Task
	AddNewTask(newTask Task)
	GetTaskById(taskId string) (Task, error)
	DeleteTaskById(taskId string) error
	UpdateTaskById(updatedTask Task) error
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
func (repo *taskRepositoryT) GetTaskById(taskId string) (Task, error) {
	var filtered []Task
	for i := 0; i < len(repo.tasks); i++ {
		task := repo.tasks[i]
		if task.Id == taskId {
			filtered = append(filtered, task)
		}
	}
	if len(filtered) == 1 {
		return filtered[0], nil
	} else {
		return Task{}, errors.New("No such task with id: " + taskId)
	}
}
func (repo *taskRepositoryT) DeleteTaskById(taskId string) error {
	_, err := repo.GetTaskById(taskId)
	if err != nil {
		return err
	}
	var newTasks []Task
	for i := 0; i < len(repo.tasks); i++ {
		currentTask := repo.tasks[i]
		if currentTask.Id != taskId {
			newTasks = append(newTasks, currentTask)
		}
	}
	repo.tasks = newTasks
	return nil
}
func (repo *taskRepositoryT) UpdateTaskById(updatedTask Task) error {
	var isUpdated = false
	for i := range repo.tasks {
		if repo.tasks[i].Id == updatedTask.Id {
			repo.tasks[i].Title = updatedTask.Title
			repo.tasks[i].Description = updatedTask.Description
			repo.tasks[i].Status = updatedTask.Status
			isUpdated = true
		}
	}
	if !isUpdated {
		return errors.New("No such task")
	}
	return nil
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
