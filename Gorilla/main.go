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

	r.Use(func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	log.Println("Servidor gorilla en ejecución en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}