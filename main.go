package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type Task struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Status string `json:"status"`
}

var tasks = []Task{}
var nextID = 1

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var task Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = nextID
    nextID++
    task.Status = "pending"
    tasks = append(tasks, task)
    json.NewEncoder(w).Encode(task)
}

// Update task status to completed
func updateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    for i, t := range tasks {
        if t.ID == id {
            tasks[i].Status = "completed"
            json.NewEncoder(w).Encode(tasks[i])
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Task not found")
}

// Delete a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    for i, t := range tasks {
        if t.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Task not found")
}

func main() {
    http.HandleFunc("/tasks", getTasks)          // GET
    http.HandleFunc("/task", createTask)         // POST
    http.HandleFunc("/task/update", updateTask)  // PUT
    http.HandleFunc("/task/delete", deleteTask)  // DELETE

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
