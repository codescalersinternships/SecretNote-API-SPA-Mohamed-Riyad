package main

import (
	"fmt"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/database"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/repository"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/server"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	r.Handle("/create-note", newServer.ValidateToken(http.HandlerFunc(newServer.CreateNote))).Methods("POST")
	r.Handle("/get-note/{noteId}", newServer.ValidateToken(http.HandlerFunc(newServer.GetNote))).Methods("GET")
	r.Handle("/get-all-notes", newServer.ValidateToken(http.HandlerFunc(newServer.GetAllNotes))).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println(err)
	}
}
