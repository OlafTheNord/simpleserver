package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our website!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jd := User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}
	jdjs, err := json.Marshal(jd)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jdjs)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userHandler)

	fmt.Println("server is starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("port is not avaible", err)
	}
}
