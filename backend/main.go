package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	r.Get("/api/tasks", getTasksFromJson)
	r.Post("/api/tasks", createTask)
	r.Put("/api/tasks/{taskId}", updateTasks)
	r.Delete("/api/tasks/{taskId}", deleteTask)
	printMonkeyLogo()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		http.ListenAndServe("localhost:9000", r)
	}()

	<-done
	fmt.Println("Server stopped.")
}
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func printMonkeyLogo() {
	fmt.Println("🐵🐒🙈🙉🙊 Welcome to Monkey Todo App! 🙊🙉🙈🐒🐵")
}
