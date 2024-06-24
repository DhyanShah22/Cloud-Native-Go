package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {

}

func (b Book) toJSON() []byte {
	return nil
}