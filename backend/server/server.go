package server

import (
	"encoding/json"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/models"
	"github.com/codescalersinternships/SecretNote-API-SPA-Mohamed-Riyad/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Server struct {
	userRepository repository.UserRepository
	noteRepository repository.NoteRepository
}

func NewServer(userRepository repository.UserRepository, noteRepository repository.NoteRepository) Server {
	return Server{
		userRepository: userRepository,
		noteRepository: noteRepository,
	}
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.userRepository.Create(&user)
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ID)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedUser, err := s.userRepository.GetByUsername(user.UserName)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if user.Password != storedUser.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ID)
}

func updateNoteViewCount(note *models.Note, s *Server) error {
	note.ViewCount = note.ViewCount + 1
	err := s.noteRepository.Update(note)
	if err != nil {
		return err
	}
	if note.ViewCount > 5 {
		err = s.noteRepository.Delete(note.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Server) CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.noteRepository.Create(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (s *Server) GetNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	noteIDStr := vars["noteId"]
	noteID, err := strconv.ParseUint(noteIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note, err := s.noteRepository.GetByID(uint(noteID))
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	err = updateNoteViewCount(note, s)
	if err != nil {
		http.Error(w, "Note cant be updated", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (s *Server) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	userIDStr := vars["userId"]
	userId, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	notes, err := s.noteRepository.GetAllByUserID(uint(userId))
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}
