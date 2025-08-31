# Todolist Go Api

Simple and friendly api-server for managing simple tasks

## How to run
`make run`

OR

`docker compose up --build --force-recreate -d`

OR

```bash
go mod download
go build -ldflags="-w -s" -o todolist-api ./cmd
./todolist-api
```

Api-server will be run on `http://localhost:9090`