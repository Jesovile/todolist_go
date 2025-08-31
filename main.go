package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webHelloWorld/tasks"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	router, startServer := InitApiServer()
	if router != nil {
		tasks.SetTasksApi(router)
		// ... add routes here
		startServer("localhost:9090")
	}
}
