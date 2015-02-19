package main

import (
	"github.com/ecopony/gamedayapi"
	"bytes"
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
	http.HandleFunc("/favicon.ico", serveFavicon)
	http.HandleFunc("/game", serveGame)
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving game")
	teamCode := r.FormValue("teamCode")
	dateString := r.FormValue("date")
	date, err := time.Parse("2006-01-02", dateString)

	if len(teamCode) == 0 || len(dateString) == 0 || err != nil {
		http.Error(w, "Three-character team code and date (yyyy-mm-dd) required", http.StatusBadRequest)
		return
	}

	game, err := gamedayapi.GameFor(teamCode, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, gameJson(game))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving template")
	teams := gamedayapi.TeamsForYear(2014)
	fp := path.Join("templates", r.URL.Path)
	tmpl, _ := template.ParseFiles(fp)
	tmpl.ExecuteTemplate(w, "index", teams)
}

func gameJson(game *gamedayapi.Game) string {
	var buffer bytes.Buffer
	buffer.WriteString(`{ "game": `)
	gameJson, _ := json.Marshal(game)
	boxscoreJson, _ := json.Marshal(game.Boxscore())
	linescoreJson, _ := json.Marshal(game.Boxscore().Linescore)
	hitchartJson, _ := json.Marshal(game.HitChart())
	buffer.WriteString(string(gameJson))
	buffer.WriteString(`, "boxscore": `)
	buffer.WriteString(string(boxscoreJson))
	buffer.WriteString(`, "linescore": `)
	buffer.WriteString(string(linescoreJson))
	buffer.WriteString(`, "hitchart": `)
	buffer.WriteString(string(hitchartJson))
	buffer.WriteString("}")
	return buffer.String()
}

func serveFavicon(w http.ResponseWriter, r *http.Request) {}
