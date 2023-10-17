package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           600,
	})

	r.Use(middleware.Logger)
	r.Use(JSONMiddleware)

	printMonkeyLogo()
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Get("/api/tasks", getTasksFromJson)
	r.Post("/api/tasks", createTask)
	r.Put("/api/tasks/{taskId}", updateTasks)
	r.Delete("/api/tasks/{taskId}", deleteTask)

	// Wrap router with the CORS handler
	handler := c.Handler(r)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Start server with the CORS-wrapped router
		http.ListenAndServe("localhost:9000", handler)
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
