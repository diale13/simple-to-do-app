package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(JSONMiddleware)
	r.Get("/",
		func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte("Hello World!"))
		})
	r.Get("/api/tasks", tasks)
	r.Post("/api/tasks", createTask)
	r.Put("/api/tasks/{taskId}", updateTasks)
	http.ListenAndServe("localhost:9000", r)
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

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
	body := UpdateTaskBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		handleError(w, http.StatusBadRequest, "Error reading request", err)
		return
	}

	tasks, err := readTasksFromJSON()
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error reading JSON file", err)
		return
	}

	tasksID, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error converting task ID from string into integer", err)
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
		handleError(w, http.StatusNotFound, "Task with ID not found", nil)
		return
	}

	if err := writeTasksToJSON(tasks); err != nil {
		handleError(w, http.StatusInternalServerError, "Error writing tasks to JSON file", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
