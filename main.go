package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func getTodos(res http.ResponseWriter, req *http.Request) {

}

func getTodo(res http.ResponseWriter, req *http.Request) {

}

func createTodo(res http.ResponseWriter, req *http.Request) {

}

func updateTodo(res http.ResponseWriter, req *http.Request) {

}

func deleteTodo(res http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	database, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Database Connection Failed")
	} else {
		fmt.Println("Database Connected Sucessfully")
	}
	database.AutoMigrate(&Todo{})
	router.HandleFunc("/api/todo", getTodos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/api/todo/new", createTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")
	fmt.Println("Sever started at Port 8000")
	http.ListenAndServe(":8000", router)
}
