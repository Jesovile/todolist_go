# Todolist Go Api

Simple and friendly api-server for managing simple tasks

## How to run

`make run`

OR

`docker compose up --build --force-recreate -d`

OR

Prerequisites: go should be installed on your machine

```bash
go mod download
go build -ldflags="-w -s" -o todolist-api ./cmd
./todolist-api
```

Api-server will be run on `http://localhost:9090`

## Api endpoints

```go
internal/tasks/api.go

router.GET("/task", getAllTasks)
router.POST("task", addNewTask)
router.DELETE("/task/:id", deleteTaskById)
router.PATCH("/task/:id", updateTask)
```

## Data model
```go
internal/tasks/data.go

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
```

## Postman collection for API-testing
`docs/Go-todolist.postman_collection.json`