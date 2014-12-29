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
	http.HandleFunc("/game", serveGame)
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveGame(w http.ResponseWriter, r *http.Request) {
	date, _ := time.Parse("2006-01-02", "2014-06-22")
	game, _ := gamedayapi.GameFor("sea", date)
	var buffer bytes.Buffer
	buffer.WriteString(`{ "game": `)
	gameJson, _ := json.Marshal(game)
	boxscoreJson, _ := json.Marshal(game.Boxscore())
	linescoreJson, _ := json.Marshal(game.Boxscore().Linescores)
	hitchartJson, _ := json.Marshal(game.HitChart())
	buffer.WriteString(string(gameJson))
	buffer.WriteString(`, "boxscore": `)
	buffer.WriteString(string(boxscoreJson))
	buffer.WriteString(`, "linescore": `)
	buffer.WriteString(string(linescoreJson))
	buffer.WriteString(`, "hitchart": `)
	buffer.WriteString(string(hitchartJson))
	buffer.WriteString("}")
	fmt.Fprintln(w, buffer.String())
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	teams := gamedayapi.TeamsForYear(2014)
	fp := path.Join("templates", r.URL.Path)
	tmpl, _ := template.ParseFiles(fp)
	tmpl.ExecuteTemplate(w, "index", teams)
}
