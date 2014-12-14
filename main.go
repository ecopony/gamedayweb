package main

import (
	"fmt"
	"github.com/ecopony/gamedayapi"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	teams := gamedayapi.TeamsForYear(2014)
	for _, team := range teams {
		fmt.Fprintln(w, team)
	}

}

func main() {
	http.HandleFunc("/", defaultHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
