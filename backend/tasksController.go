package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func createTask(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a CreateTaskBody struct
	var body CreateTaskBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleError(w, http.StatusInternalServerError, "Error decoding request body", err)
		return
	}

	// Read the existing tasks from the JSON file
	tasks, err := readTasksFromJSON()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error reading tasks from JSON file", err)
		return
	}

	// Check if the task name is empty
	if body.Name == "" {
		handleError(w, http.StatusBadRequest, "Task name cannot be empty", nil)
		return
	}

	// Create a new task and add it to the end of the list
	newTask := Task{
		Name: body.Name,
		Done: false,
		ID:   len(tasks) + 1,
	}
	tasks = append(tasks, newTask)

	// Write the updated task list back to the JSON file
	if err := writeTasksToJSON(tasks); err != nil {
		handleError(w, http.StatusInternalServerError, "Error writing tasks to JSON file", err)
		return
	}

	// Return the newly created task as a response
	respondWithJSON(w, http.StatusCreated, newTask)
}

func readTasksFromJSON() ([]Task, error) {
	jsonFile, err := os.Open(JSONFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var tasks []Task
	if err := json.NewDecoder(jsonFile).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasksToJSON(tasks []Task) error {
	j, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	if err := os.WriteFile(JSONFilePath, j, 0755); err != nil {
		return err
	}

	return nil
}

func handleError(w http.ResponseWriter, status int, message string, err error) {
	w.WriteHeader(status)
	log.Printf("%s: %v", message, err)
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		handleError(w, http.StatusInternalServerError, "Error encoding JSON response", err)
	}
}

func updateTasks(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into an UpdateTaskBody struct
	body := UpdateTaskBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleError(w, http.StatusBadRequest, "Error reading request body", err)
		return
	}

	// Read the existing tasks from the JSON file
	tasks, err := readTasksFromJSON()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error reading tasks from JSON file", err)
		return
	}

	// Convert the task ID from a string to an integer
	tasksID, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		handleError(w, http.StatusBadRequest, "Error converting task ID from string to integer", err)
		return
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == tasksID {
			if body.Name != nil {
				task.Name = *body.Name
			}
			if body.Done != nil {
				task.Done = *body.Done
			}
			tasks[i] = task
			taskFound = true
			break
		}
	}

	if !taskFound {
		handleError(w, http.StatusNotFound, "Task with the specified ID not found", nil)
		return
	}

	// Write the updated task list back to the JSON file
	if err := writeTasksToJSON(tasks); err != nil {
		handleError(w, http.StatusInternalServerError, "Error writing tasks to JSON file", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Read the existing tasks from the JSON file
	tasks, err := readTasksFromJSON()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error reading tasks from JSON file", err)
		return
	}

	// Convert the task ID from a string to an integer
	tasksID, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		handleError(w, http.StatusBadRequest, "Error converting task ID from string to integer", err)
		return
	}

	// Variable to track if the task with the specified ID is found
	taskFound := false

	// Loop through the tasks to find the one with the matching ID
	for i, task := range tasks {
		if task.ID == tasksID {
			// Remove the task from the slice.
			tasks = append(tasks[:i], tasks[i+1:]...)
			taskFound = true
			break
		}
	}
	if !taskFound {
		handleError(w, http.StatusNotFound, "Task with the specified ID not found", nil)
		return
	}

	// Write the updated task list back to the JSON file
	if err := writeTasksToJSON(tasks); err != nil {
		handleError(w, http.StatusInternalServerError, "Error writing tasks to JSON file", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
