package main

import (
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

func todos(res http.ResponseWriter, req *http.Request) {

}

func todo(res http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Database Connection Failed")
	} else {
		fmt.Println("Database Connected Sucessfully")
	}
	database.AutoMigrate(&Todo{})
	router.HandleFunc("/api/todo", todos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", todo).Methods("GET")
	fmt.Println("Sever started at Port 8000")
	http.ListenAndServe(":8000", router)
}
