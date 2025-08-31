package tasks

import "fmt"

func ApiMain() {
	tasks := TaskRepository.GetAllTasks()
	fmt.Println(tasks)
}
