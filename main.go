package main

import (
	"github.com/ecopony/gamedayapi"
	"log"
	"net/http"
	"path"
	"html/template"
	"fmt"
	"time"
	"encoding/json"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/game", serveGame)
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveGame(w http.ResponseWriter, r *http.Request) {
	date, _ := time.Parse("2006-01-02", "2014-06-22")
	game, _ := gamedayapi.GameFor("sea", date)
	gameJson, _ := json.Marshal(game)
	fmt.Fprintln(w, string(gameJson))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	teams := gamedayapi.TeamsForYear(2014)
	fp := path.Join("templates", r.URL.Path)
	tmpl, _ := template.ParseFiles(fp)
	tmpl.ExecuteTemplate(w, "index", teams)
}
