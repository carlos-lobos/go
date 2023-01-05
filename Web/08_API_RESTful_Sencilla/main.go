package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

var noteStore = make(map[string]Note)

var id int

// [GET] /api/notes
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(notes)
	if err != nil {
		log.Println("[ERROR] No se pudo codificar las notas en json.")
		j, err = json.Marshal(err)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(j)
}

// [POST] /api/notes
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println("[ERROR] No se pudo decodificar la nota desde json.")
		j, err := json.Marshal(err)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
	} else {
		note.CreatedAt = time.Now()
		id++
		k := strconv.Itoa(id)
		noteStore[k] = note

		j, err := json.Marshal(note)
		if err != nil {
			log.Println("[ERROR] No se pudo re-codificar la nota en json.")
			// j, err = json.Marshal(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}

		w.Write(j)
	}

}

// [PUT] /api/notes/{id}
func ModifyNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	// w.Header().Set("Content-Type", "application/json")

	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		log.Println("[ERROR] No se pudo decodificar la nota desde json.")
		// j, err := json.Marshal(err)
		// w.WriteHeader(http.StatusBadRequest)
	} else {
		if note, ok := noteStore[k]; ok {
			noteUpdate.CreatedAt = note.CreatedAt
			delete(noteStore, k)
			noteStore[k] = noteUpdate
		} else {
			log.Printf("[ERROR] No se encontro el id de la nota %s.", k)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// [DELETE] /api/notes/{id}
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("[ERROR] No se encontro el id de la nota %s.", k)
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNotes).Methods("GET")
	r.HandleFunc("/api/notes", CreateNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", ModifyNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNote).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening http://localhost:8080 ...")
	log.Fatal(server.ListenAndServe())
}
