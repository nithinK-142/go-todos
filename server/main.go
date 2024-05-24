package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Todo struct {
	ID          int    `json:"id"`
	IsDone      bool   `json:"isDone"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var todos = []Todo{
	{ID: 1, IsDone: false, Name: "Learn Go", Description: "Complete the Go programming language tutorial"},
	{ID: 2, IsDone: true, Name: "Buy groceries", Description: "Get milk, eggs, and bread from the store"},
	{ID: 3, IsDone: false, Name: "Clean the house", Description: "Vacuum, mop, and dust the living room"},
}

var nextTodoID = 4

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Here we go!!")
	})
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", editTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router)))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newTodo.ID = nextTodoID
	nextTodoID++
	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func editTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	var editedTodo Todo
	err = json.NewDecoder(r.Body).Decode(&editedTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = editedTodo
			todos[i].ID = id
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Todo edited successfully")
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Todo deleted successfully")
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}
