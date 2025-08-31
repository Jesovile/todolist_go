package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webHelloWorld/internal/tasks"
)

func InitApiServer() (*gin.Engine, func(addr string)) {
	router := gin.Default()
	startServer := func(addr string) {
		err := router.Run(addr)
		if err != nil {
			fmt.Println("Server starting failed", err)
		}
	}
	return router, startServer
}

func main() {
	router, startServer := InitApiServer()
	if router != nil {
		tasks.SetTasksApi(router)
		// ... add routes here
		startServer(":9090")
	}
}
