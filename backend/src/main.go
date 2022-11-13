package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Category struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}

type Categories []Category

var categories = Categories{
	Category{Name: "Restaurants", Description: "Places we like to eat", Content: "A place to talk about food", CreatedAt: time.Now()},
	Category{Name: "Date Nights", Description: "Places for fun couple activities", Content: "A place to talk about activities", CreatedAt: time.Now()},
	Category{Name: "Movies", Description: "Best Films", Content: "A place to talk about movies", CreatedAt: time.Now()},
}

func allCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Categories Endpoint")
	json.NewEncoder(w).Encode(categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Create Category Endpoint")
	var newCategory Category
	json.NewDecoder(r.Body).Decode(&newCategory)
	categories = append(categories, newCategory)
	json.NewEncoder(w).Encode(newCategory)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete Category Endpoint")
	vars := mux.Vars(r)
	name := vars["name"]
	for index, category := range categories {
		if category.Name == name {
			categories = append(categories[:index], categories[index+1:]...)
		}
	}
}

func editCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Edit Category Endpoint")
	vars := mux.Vars(r)
	name := vars["name"]
	for index, category := range categories {
		if category.Name == name {
			categories = append(categories[:index], categories[index+1:]...)
			var newCategory Category
			json.NewDecoder(r.Body).Decode(&newCategory)
			categories = append(categories, newCategory)
			json.NewEncoder(w).Encode(newCategory)
		}
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", test)
	myRouter.HandleFunc("/categories", allCategories).Methods("GET")
	myRouter.HandleFunc("/categories", createCategory).Methods("POST")
	myRouter.HandleFunc("/categories/{name}", deleteCategory).Methods("DELETE")
	myRouter.HandleFunc("/categories/{name}", editCategory).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
