package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
