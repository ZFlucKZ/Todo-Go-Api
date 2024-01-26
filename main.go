package main

// Create a server that listens on port 8080 and which has a handler for the route "/".
// Have the handler return the response "Hello world"

import (
	"fmt"
	"net/http"

	"github.com/ZFlucKZ/Todo-Go-Api/db"
	"github.com/ZFlucKZ/Todo-Go-Api/routes"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	db.Init()

	routes.RegisterTodoRoutes()

	http.ListenAndServe(":8080", nil)
}
