package main

import (
	"fmt"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/database"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/repository"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/server"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db, err := database.InitializeDB()
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}
	noteRepository := repository.NewNoteRepository(db)
	userRepository := repository.NewUserRepository(db)
	newServer := server.NewServer(userRepository, noteRepository)

	r := mux.NewRouter()
	r.HandleFunc("/signup", newServer.SignUp).Methods("POST")
	r.HandleFunc("/signin", newServer.SignIn).Methods("POST")
	r.HandleFunc("/createNote", newServer.CreateNote).Methods("POST")
	r.HandleFunc("/getNote/{noteId}", newServer.GetNote).Methods("GET")
	r.HandleFunc("/getAllNotes/{userId}", newServer.GetAllNotes).Methods("GET")

	fmt.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
