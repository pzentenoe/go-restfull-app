package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"fmt"
)

func DELETEUsers(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "DELETE")
}
func PUTUsers(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "PUT")
}
func POSTUsers(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "POST")
}
func GetUsers(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "GET")
}

func main() {

	router := mux.NewRouter().StrictSlash(false) //Hace que las rutas sean distintas aunque terminen con slash
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users", POSTUsers).Methods("POST")
	router.HandleFunc("/api/users", PUTUsers).Methods("PUT")
	router.HandleFunc("/api/users", DELETEUsers).Methods("DELETE")

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
