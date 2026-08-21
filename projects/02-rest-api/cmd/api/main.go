// Command api is Project 2: an in-memory REST API for tasks (Postgres
// arrives in Project 4). This file is wiring only, per Stage 6's own
// layered structure — config/deps/router/server, nothing else.
package main

import (
	"log"
	"net/http"

	"github.com/YOUR_USERNAME/go-backend-learning/projects/02-rest-api/internal/task"
)

func main() {
	repo := task.NewRepository()
	svc := task.NewService(repo)
	handler := task.NewHandler(svc)
	mux := task.NewMux(handler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
