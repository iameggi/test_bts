package main

import (
	"todolist/router"
)

func main() {
	// Initialize the router
	r := router.SetupRouter()

	// Start server
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
