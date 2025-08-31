package tasks

// Tasks Repository Types

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskRepositoryI interface {
	GetAllTasks() []Task
}

type taskRepositoryT struct {
	tasks []Task
}

// todo: delete this fake, for test only
func getStubData(count int) []Task {
	var result []Task
	for i := 1; i <= count; i++ {
		var stringIndex = string(i)
		taskItem := Task{stringIndex, "Task title " + stringIndex, "Task description for Task " + stringIndex, "todo"}
		result = append(result, taskItem)
	}
	return result
}

// GetAllTasks Tasks Repository Implementation
func (repo taskRepositoryT) GetAllTasks() []Task {
	return repo.tasks
}

// TaskRepository Singleton export
var TaskRepository = taskRepositoryT{
	tasks: getStubData(15),
}
