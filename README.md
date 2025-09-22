# Go To-Do REST API

A RESTful API built in **Go** to manage tasks with full **CRUD** functionality (Create, Read, Update, Delete). Ideal for beginners to learn Go, HTTP handling, and JSON data management.

## Features
- Create new tasks
- Retrieve all tasks
- Update task status to completed
- Delete tasks
- JSON input/output
- Lightweight in-memory storage

## Technologies
- Go
- net/http
- JSON

## API Endpoints

| Method | Endpoint           | Description                  |
|--------|------------------|------------------------------|
| GET    | /tasks            | Get all tasks               |
| POST   | /task             | Create a new task           |
| PUT    | /task/update      | Update a task status        |
| DELETE | /task/delete      | Delete a task               |

## How to Run Locally

1. Install Go: [https://go.dev/dl/](https://go.dev/dl/)  
2. Clone this repository:

```bash
git clone https://github.com/ritikkumar55/go-todo-api.git
cd go-todo-api
go run main.go
