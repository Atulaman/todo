package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type task struct {
	Id   int    `json:"id"`
	Desc string `json:"desc"`
}

var tasks []task
var taskId int

func Add(w http.ResponseWriter, r *http.Request) {
	var newtask task
	err := json.NewDecoder(r.Body).Decode(&newtask)
	if err != nil || newtask.Desc == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	taskId++
	newtask.Id = taskId
	tasks = append(tasks, newtask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newtask)
}
func List(w http.ResponseWriter, r *http.Request) {
	if len(tasks) == 0 {
		http.Error(w, "No tasks found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	//json.NewEncoder(w).Encode(map[string]string{"message": "Tasks found successfully"})
}
func Update(w http.ResponseWriter, r *http.Request) {
	var modtask task
	err := json.NewDecoder(r.Body).Decode(&modtask)
	if err != nil || modtask.Desc == "" || modtask.Id == 0 {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	for i, t := range tasks {
		if t.Id == modtask.Id {
			tasks[i].Desc = modtask.Desc
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(modtask)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	var deltask task
	err := json.NewDecoder(r.Body).Decode(&deltask)
	if err != nil || deltask.Id == 0 {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	for i, t := range tasks {
		if t.Id == deltask.Id {
			tmp1 := tasks[:i]
			tmp2 := tasks[i+1:]
			tasks = append(tmp1, tmp2...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to My_TODO_API")
	})
	http.HandleFunc("/tasks/Add", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			Add(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			List(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/tasks/update", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPut {
			Update(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/tasks/delete", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodDelete {
			Delete(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})
	// var t = task{
	// 	Id:   1,
	// 	Desc: "Task 1",
	// }
	// tasks = append(tasks, t)
	// var p = task{
	// 	Id:   2,
	// 	Desc: "Task 2",
	// }
	// tasks = append(tasks, p)
	fmt.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
