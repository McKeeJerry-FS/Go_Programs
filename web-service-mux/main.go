package main

import(
	"encoding/json"
	//"fmt"
	//"errors"
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

// type IdNotFoundError struct{
// 	ID string
// }

// func (e IdNotFoundError) Error() string {
// 	return fmt.Sprintf("Book with ID %d not found", e.ID)
// }

// Function for returning all books
func GetBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

// Function for returning a single book with a matching ID
func GetBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
	
	// if err != nil {
	// 	if _, ok := err.(IdNotFoundError); ok {
	// 		http.Error(w, err.Error(), http.StatusNotFound)
	// 		return
	// 	}
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

// Function for creating a book
func CreateBook(w http.ResponseWriter, r *http.Request){
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Function for updating a book
func UpdateBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Function for deleting a book
func DeleteBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}



func main() {
	router := mux.NewRouter()

	books =  append(books, Book{ID: "1", Title: "Book One", Author: "John Doe", Publisher: &Company{Name: "Publisher One", Address: "Address One"}})
	books =  append(books, Book{ID: "2", Title: "Book Two", Author: "Jane Smith", Publisher: &Company{Name: "Publisher Two", Address: "Address Two"}})


	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

