package todo

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// Todo Struct
type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

//Routes for all the API endpoints
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo)
	router.Delete("/{todoID}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)

	return router
}

//GetATodo fetches a Todo using the todoId from URL params
func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "Hello",
		Body:  "Hello Body",
	}
	render.JSON(w, r, todos) // Return a demo response
}

//DeleteTodo just deletes a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Todo Successfully"
	render.JSON(w, r, response) // Return a demo response
}

//CreateTodo just creates a todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Todo Successfully"
	render.JSON(w, r, response) //Return a demo response
}

//GetAllTodos just fetches all todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug:  "slug",
			Title: "Hello",
			Body:  "Hello Body",
		},
		{
			Slug:  "slug",
			Title: "Hello",
			Body:  "Hello Body",
		},
	}
	render.JSON(w, r, todos)
}
