package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"fmt"
	"encoding/json"
	"strconv"
)

var noteStore = make(map[string]Note)

var id int

func GetNoteHandler(w http.ResponseWriter, request *http.Request) {
	var notes [] Note
	for _, value := range noteStore {
		notes = append(notes, value)
	}
	w.Header().Set("Content-Type", "application/json")
	j, error := json.Marshal(notes)
	if error != nil {
		panic(error)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func PostNoteHandler(w http.ResponseWriter, request *http.Request) {
	var note Note
	//Decodificar json a Objeto
	err := json.NewDecoder(request.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedAt = time.Now()
	id ++
	k := strconv.Itoa(id)
	noteStore[k] = note

	w.Header().Set("Content-Type", "application/json")
	j, error := json.Marshal(note)
	if error != nil {
		panic(error)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}
func PutNoteHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "PUT")
}
func DeleteNoteHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "DELETE")
}

type Note struct {
	Title       string `json:"title"` //Notacion para decir como se llamara el campo al convertirlo en json
	Description string `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {

	router := mux.NewRouter().StrictSlash(false) //Hace que las rutas sean distintas aunque terminen con slash
	router.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	router.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	router.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	router.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //maximo megabytes del header, operador shift multiplica por 2 y eleva 20 veces y devuelve en bytes
	}

	log.Println("Listening")
	server.ListenAndServe()

}
