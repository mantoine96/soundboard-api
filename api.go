package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiResponse struct {
	ErrorMessage string      `json:"error,omitempty"`
	ID           string      `json:"id,omitempty"`
	Result       interface{} `json:"result,omitempty"`
}

type apiHandlerFunc func(http.ResponseWriter, map[string]string, *json.Encoder, *json.Decoder)

func test(router *mux.Router) {
	log.Println("Test func")

}
