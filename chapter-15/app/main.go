package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(w http.ResponseWriter, req *http.Request) {
	write(w, "Hello, web")
}
func frenchHandler(w http.ResponseWriter, req *http.Request) {
	write(w, "Salut web")
}

func hindiHandler(w http.ResponseWriter, req *http.Request) {
	write(w, "Namaste, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
