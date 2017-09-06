package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"soundboard-api/soundboard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	router := soundboard.NewRouter()
	request, _ := http.NewRequest("GET", "/status", nil)
	response := httptest.NewRecorder()
	//expectedResponse, _ := json.Marshal(map[string]string{"status": "ok"})
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "HTTP200 is expected")
	var responseJSON map[string]string
	err := json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		t.Errorf("Error when decoding JSON")
	}
	assert.Equal(t, map[string]string{"status": "ok"}, responseJSON, "Status should be OK")
}
func TestGetSounds(t *testing.T) {

	t.Errorf("Not implemented")
}
func TestAddSound(t *testing.T) {
	t.Errorf("Not implemented")
}

func TestUpdateSound(t *testing.T) {
	t.Errorf("Not implemented")
}
func TestDeleteSound(t *testing.T) {
	t.Errorf("Not implemented")
}
