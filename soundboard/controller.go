package soundboard

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Create Controller type
type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	sounds := c.Repository.GetSounds() // Get list of sounds
	log.Println(sounds)
	data, _ := json.Marshal(sounds)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddSound POST /
func (c *Controller) AddSound(w http.ResponseWriter, r *http.Request) {
	var sound Sound
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // Read body of request. 1048576 is default max request size
	if err != nil {
		log.Fatalln("Error AddSound", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddSound", err)
	}
	fmt.Println(body)
	if err := json.Unmarshal(body, &sound); err != nil { // Unmarshall body contents as a type Sound (JSON decode)
		w.WriteHeader(422) // Can't process
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddSound unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.AddSound(sound)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateSound PUT /
func (c *Controller) UpdateSound(w http.ResponseWriter, r *http.Request) {
	var sound Sound
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // See Line 33
	if err != nil {
		log.Fatalln("Error UpdateSound", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateSound", err)
	}
	if err := json.Unmarshal(body, &sound); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // Can't process entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateSound unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.UpdateSound(sound) // Update sound in DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteSound DELETE /
func (c *Controller) DeleteSound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.Repository.DeleteSound(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}

// Status GET /
func (c *Controller) GetStatus(w http.ResponseWriter, r *http.Request) {

	status := map[string]string{"status": "ok"}
	data, err := json.Marshal(status)
	if err != nil {
		log.Fatalln("Error status", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
