package main

import (
	"database/sql"
	"time"

	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
	"./server"
)

func main() {
	db, _ := sql.Open("sqlite3", "./cricket.db")
	defer db.Close()

	c := colly.NewCollector()

	c.OnHTML("div.cscore_link cscore_link--button", func(e *colly.HTMLElement) {
		teams := e.ChildText("div.cscore_info-overview > div.cscore_name cscore_name--long")
		score := e.ChildText("div.cscore_notes_game > div.cscore_score")
		db.Exec("INSERT INTO cricket_scores(teams, score) VALUES(?,?)", teams, score)
	})

	go func() {
		for {
			c.Visit("https://www.espn.in/cricket/scores")
			time.Sleep(2 * time.Second)
		}
	}()

	StartServer(db)
}
