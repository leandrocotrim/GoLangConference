package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("/users/{id}", getUserHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Dia3 - Home Go!")
}

func erroHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func strErroHandler(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusInternalServerError)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "db/dia3.db")
	if err != nil {
		erroHandler(w, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT Id, Name, Email FROM USERS")
	if err != nil {
		erroHandler(w, err)
		return
	}
	defer rows.Close()
	
	users := []*user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
			erroHandler(w, err)
			return
		}
		users = append(users, &u)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		erroHandler(w, err)
		return	
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "db/dia3.db")
	if err != nil {
		erroHandler(w, err)
		return
	}
	defer db.Close()

	var id = r.PathValue("id")
	if id == "" {
		strErroHandler(w, "Missing id parameter")
		return
	}

	var u user
	if err := db.QueryRow("SELECT Id, Name, Email FROM USERS WHERE Id = ?", id).Scan(&u.Id, &u.Name, &u.Email); err != nil {
		erroHandler(w, err)
		return
	}
	
	if err := json.NewEncoder(w).Encode(u); err != nil {
		erroHandler(w, err)
		return	
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "db/dia3.db")
	if err != nil {
		erroHandler(w, err)
		return
	}
	defer db.Close()

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		erroHandler(w, err)
		return	
	}

	if _, err := db.Exec("INSERT INTO USERS (Name, Email) VALUES(?, ?)", u.Name, u.Email); err != nil {
		erroHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
