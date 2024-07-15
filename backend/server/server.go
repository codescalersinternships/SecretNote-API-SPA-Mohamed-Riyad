package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Note struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ViewCount int    `json:"view_count"`
}

type user struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

var db *sql.DB

func signUp(w http.ResponseWriter, r *http.Request) {
	var user user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO users (user_name, password) VALUES (?, ?)", user.UserName, user.Password)
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func signIn(w http.ResponseWriter, r *http.Request) {

	var user user

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE user_name = ?", user.UserName).Scan(&storedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if user.Password != storedPassword {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateNoteViewCount(note Note, id int) error {
	note.ViewCount++

	if note.ViewCount > 5 {

		delStmt, err := db.Prepare("DELETE FROM notes WHERE id = ?")
		if err != nil {
			return err
		}
		defer delStmt.Close()

		_, err = delStmt.Exec(id)
		if err != nil {
			return err
		}
		return fmt.Errorf("note deleted due to exceeding view limit")
	} else {
		updateStmt, err := db.Prepare("UPDATE notes SET view_count = ? WHERE id = ?")
		if err != nil {
			return err
		}
		defer updateStmt.Close()

		_, err = updateStmt.Exec(note.ViewCount, id)
		if err != nil {
			return err
		}
	}
	return nil
}
func createNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO notes(user_id, title, content) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(note.UserID, note.Title, note.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	note.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("SELECT id, user_id, title, content, view_count FROM notes WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var note Note
	err = stmt.QueryRow(id).Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.ViewCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Note not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = updateNoteViewCount(note, note.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func getAllNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("SELECT id, user_id, title, content, view_count FROM notes WHERE user_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.ViewCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = updateNoteViewCount(note, note.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func main() {

	db, err := sql.Open("sqlite3", "./dataBase.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_name TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );
    `)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS notes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL references users(id),
            title TEXT NOT NULL ,
            content TEXT NOT NULL,
            view_count INTEGER DEFAULT 0
        );
    `)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/signin", signIn)
	http.HandleFunc("/createNote", createNote)
	http.HandleFunc("/getNote", getNote)
	http.HandleFunc("/getAllNotes", getAllNotes)

	fmt.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
