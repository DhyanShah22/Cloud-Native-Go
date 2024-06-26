package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "D. A. Shah", ISBN: "22092210"}
	json := book.ToJSON()

	expectedJSON := `{"title":"Cloud Native Go","author":"D. A. Shah","isbn":"22092210"}`
	assert.JSONEq(t, expectedJSON, string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	jsonData := []byte(`{"title":"Cloud Native Go","author":"D. A. Shah","isbn":"22092210"}`)
	book := FromJSON(jsonData)

	expected := Book{Title: "Cloud Native Go", Author: "D. A. Shah", ISBN: "22092210"}
	assert.Equal(t, expected, book, "Book JSON unmarshalling wrong.")
}
