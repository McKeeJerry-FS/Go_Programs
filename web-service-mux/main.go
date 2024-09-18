package main

import(
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Defining a Book struct
type Book struct {
	ID 			string		`json:"id,omitempty"`
	Title 		string		`json:"title,omitempty"`
	Author 		string		`json:"author,omitempty"`
	Publisher 	*Company	`json:"publisher,omitempty"`
}
// Definign a Company struct
type Company struct{
	Name 		string		`json:"name,omitempty"`
	Address 	string		`json:"address,omitempty"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

