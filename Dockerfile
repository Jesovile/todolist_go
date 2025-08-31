from golang:1.25-alpine as build

workdir /build

copy go.mod go.sum ./
run go mod download

copy . .

run CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o todolist-api ./cmd

expose 9090
env DOCKERENV=1

cmd ["/build/todolist-api"]
