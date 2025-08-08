package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "¡Bienvenido al servidor Gorilla!")
	})

	log.Println("Servidor gorilla en ejecución en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}