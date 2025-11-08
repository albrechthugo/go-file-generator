package main

import (
	"html/template"
	"log"
	"net/http"
)

type GitHubProfileData struct {
	Name string
}

type GitHubProfileTemplateData struct {
	Profile GitHubProfileData
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/gh-profile.html"))

	var data = GitHubProfileTemplateData{
		Profile: GitHubProfileData{
			Name: "Hugo Albrecht",
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}