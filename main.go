package main

import (
	"github.com/ecopony/gamedayapi"
	"log"
	"net/http"
	"path"
	"html/template"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	teams := gamedayapi.TeamsForYear(2014)
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", r.URL.Path)

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", teams)
}
