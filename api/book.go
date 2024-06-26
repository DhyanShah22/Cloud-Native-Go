package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
	Description string `json:"description"`
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

var Books = []Book{
	Book{Title:"The Hitchhiker's Guide to Galaxy!", Author: "Douglas Adams", ISBN: "19761981"},
	Book{Title:"Cloud Native Go!", Author: "D. A. Shah", ISBN: "22092210"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}