package api

import (
	"encoding/json"
	"net/http"
	"io"
	"strconv"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Description string `json:"description"`
}

var books = map[string]Book{
	"19761981" : Book{Title:"The Hitchhiker's Guide to Galaxy!", Author: "Douglas Adams", ISBN: "19761981"},
	"22092210" : Book{Title:"Cloud Native Go!", Author: "D. A. Shah", ISBN: "22092210"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	switch method := r.Method; method {
		case http.MethodGet:
			books := AllBooks()
			writeJSON(w, books)
		case http.MethodPost:
			body, err := io.ReadAll(r.Body)
			if err != nil {
			 w.WriteHeader(http.StatusInternalServerError)
			}
			book := FromJSON(body)
			isbn, created := CreateBook(book)
			if created {
				w.Header().Add("Location", "/api/books/"+isbn)
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusConflict)
			}
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("Unsupported request method"))
	}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request){
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
		case http.MethodGet:
			book, found := GetBook(isbn)
			if found {
				writeJSON(w, book)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		case http.MethodPut:
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			book := FromJSON(body)
			exists := UpdateBook(isbn, book)
			if exists {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		case http.MethodDelete:
			DeleteBook(isbn)
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unsupported request method"))
	}
}

func AllBooks() []Book {
	allBooks := []Book{}
	for _, book := range books {
	  allBooks = append(allBooks, book)
	}
	return allBooks
  }

var counter int = 0
  
  func CreateBook(book Book) (string, bool) {
	counter++
	newISBN := strconv.Itoa(counter)
  
	_, exists := books[newISBN]
	if exists {
	  return "", false // Handle ISBN conflict (consider retrying with a new counter value)
	}
  
	books[newISBN] = book
	return newISBN, true
  }
  
  func writeJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
  }
  

  func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
  }
  
  func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if !exists {
	  return false
	}
	books[isbn] = book
	return true
  }
  
  func DeleteBook(isbn string) {
	delete(books, isbn)
  }

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}