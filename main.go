package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Fprintf(w, "Hello, %s!", username)
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Product ID:%s", id)
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Form")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/hello/{username}", hello)
	r.HandleFunc(`/product/{id:\d+}`, product)
	r.HandleFunc(`/form`, form).Methods("POST", "PUT")
	r.NotFoundHandler = http.HandlerFunc(handler404)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}

func handler404(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page Not Found"))
}
