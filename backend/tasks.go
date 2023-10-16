package main

import (
	"log"
	"net/http"
	"os"
)

const (
	JSONFilePath = "./tasks.json"
)

type Task struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
	Done bool   `json: "done"`
}

type CreateTaskBody struct {
	Name string `json: "name"`
}

type UpdateTaskBody struct {
	Name *string `json: "name"`
	Done *bool   `json: "done"`
}

func getTasksFromJson(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.ReadFile(JSONFilePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error while reading tasks jsonfile %v", err)
		return
	}
	w.Write(jsonFile)
}
