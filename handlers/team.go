package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type member struct {
	ID          string
	Name        string
	Description string
}

var members struct {
	Current        []member
	PastPresidents []member
	PastOther      []member
}

func init() {
	path, _ := filepath.Abs("assets/json/members.json")
	file, e := ioutil.ReadFile(path)

	if e != nil {
		log.Fatal(e)
	}

	json.Unmarshal(file, &members)
}

func team(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "team", &members)
}
