
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)


// test view all tasks
func TestGetAllTasks(t *testing.T) {
	// Initialize the tasks with 5 example tasks
	tasks = []Task{
		{ID: "1", Title: "Finish project 1", Completed: false},
		{ID: "2", Title: "Finish project 2", Completed: true},
		{ID: "3", Title: "Finish project 3", Completed: false},
		{ID: "4", Title: "Finish project 4", Completed: true},
		{ID: "5", Title: "Finish project 5", Completed: false},
	}

	// Create a new request
	req, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(t, err)

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Initialize a new router and handle the request
	router := mux.NewRouter()
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the response body and unmarshal it into a slice of tasks
	var responseTasks []Task
	err = json.Unmarshal(rr.Body.Bytes(), &responseTasks)
	assert.NoError(t, err)

	// Define the expected tasks
	expectedTasks := []Task{
		{ID: "1", Title: "Finish project 1", Completed: false},
		{ID: "2", Title: "Finish project 2", Completed: true},
		{ID: "3", Title: "Finish project 3", Completed: false},
		{ID: "4", Title: "Finish project 4", Completed: true},
		{ID: "5", Title: "Finish project 5", Completed: false},
	}

	// Compare the expected tasks with the response tasks
	assert.Equal(t, expectedTasks, responseTasks)
}


// test view specific task
func TestGetTaskByID(t *testing.T) {
	// Populate tasks with 3 test tasks
	tasks = []Task{
		{ID: "1", Title: "Task 1", Completed: false},
		{ID: "2", Title: "Task 2", Completed: true},
		{ID: "3", Title: "Task 3", Completed: false},
	}

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{taskID}", getTaskByID).Methods("GET")

	// Select the second task for testing
	taskID := "2"
	req, _ := http.NewRequest("GET", "/tasks/"+taskID, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var task Task
	err := json.Unmarshal(rr.Body.Bytes(), &task)
	assert.NoError(t, err)

	// Check if the retrieved task matches the selected task
	assert.Equal(t, taskID, task.ID)
	assert.Equal(t, "Task 2", task.Title)
	assert.True(t, task.Completed)

	// You can add more assertions here to check other properties of the task
}


// test create new task
func TestCreateTask(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", createTask).Methods("POST")

	newTask := Task{Title: "New Task", Completed: false}
	jsonBody, _ := json.Marshal(newTask)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdTask Task
	err := json.Unmarshal(rr.Body.Bytes(), &createdTask)
	assert.NoError(t, err)

	assert.Equal(t, newTask.Title, createdTask.Title)
	assert.False(t, createdTask.Completed)

	// You can add more assertions here to check other properties of the created task
}

func TestEditTask(t *testing.T) {
	// Populate tasks with 3 test tasks
	tasks = []Task{
		{ID: "1", Title: "Task 1", Completed: false},
		{ID: "2", Title: "Task 2", Completed: true},
		{ID: "3", Title: "Task 3", Completed: false},
	}

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{taskID}", editTask).Methods("PUT")

	taskID := "2"
	updatedTask := Task{ID: taskID, Title: "Updated Task", Completed: false}
	jsonBody, _ := json.Marshal(updatedTask)
	req, _ := http.NewRequest("PUT", "/tasks/"+taskID, bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var editedTask Task
	err := json.Unmarshal(rr.Body.Bytes(), &editedTask)
	assert.NoError(t, err)

	assert.Equal(t, updatedTask.Title, editedTask.Title)
	assert.False(t, editedTask.Completed)

	// You can add more assertions here to check other properties of the edited task
}

func TestDeleteTask(t *testing.T) {
	// Populate tasks with 3 test tasks
	tasks = []Task{
		{ID: "1", Title: "Task 1", Completed: false},
		{ID: "2", Title: "Task 2", Completed: true},
		{ID: "3", Title: "Task 3", Completed: false},
	}

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{taskID}", deleteTask).Methods("DELETE")

	taskID := "2"
	req, _ := http.NewRequest("DELETE", "/tasks/"+taskID, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	// Ensure that the task has been deleted
	for _, task := range tasks {
		assert.NotEqual(t, taskID, task.ID)
	}

	// You can add more assertions here to check other properties after deletion
}
