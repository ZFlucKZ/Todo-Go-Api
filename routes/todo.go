package routes

import (
	"net/http"

	"github.com/ZFlucKZ/Todo-Go-Api/models"
)

func RegisterTodoRoutes() {
	http.HandleFunc("/create", models.CreateTodo)
	http.HandleFunc("/get-all", models.GetTodos)
	http.HandleFunc("/get", models.GetTodo)
	http.HandleFunc("/edit", models.EditTodo)
	http.HandleFunc("/delete", models.DeleteTodo)
}