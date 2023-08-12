
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task

// view all tasks: GET /tasks
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// new task: POST /tasks
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	newTask.ID = fmt.Sprintf("%d", len(tasks)+1)
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}


// view task: GET /tasks/{taskID}
func getTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID := vars["taskID"]

	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.NotFound(w, r)
}


// edit task: PUT /tasks/{taskID}
func editTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID := vars["taskID"]

	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range tasks {
	if task.ID == taskID {
		updatedTask.ID = taskID // Assign the ID to updatedTask first
		tasks[i] = updatedTask  // Then update the task in the tasks slice
		json.NewEncoder(w).Encode(updatedTask)
		return
	}
}


	http.NotFound(w, r)
}

// delete task: DELETE /tasks/{taskID}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskID := vars["taskID"]

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}


func main() {
	router := mux.NewRouter()

	tasks = append(tasks, Task{ID: "1", Title: "Finish project", Completed: false})

	// REST routes defined here
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{taskID}", getTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{taskID}", editTask).Methods("PUT")
	router.HandleFunc("/tasks/{taskID}", deleteTask).Methods("DELETE")



	log.Fatal(http.ListenAndServe(":8080", router))
}