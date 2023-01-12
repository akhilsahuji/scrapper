package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Score struct {
	Teams string
	Score string
}

// StartServer starts the http server and serves the data
func StartServer(db *sql.DB) {
	http.HandleFunc("/cricket-scores", func(w http.ResponseWriter, r *http.Request) {
		rows, _ := db.Query("SELECT teams, score FROM cricket_scores")
		var scores []Score
		for rows.Next() {
			var teams, score string
			rows.Scan(&teams, &score)
			scores = append(scores, Score{teams, score})
		}
		t, _ := template.ParseFiles("template.html")
		t.Execute(w, scores)
	})
	http.ListenAndServe(":8080", nil)
}
