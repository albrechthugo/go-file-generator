package main

import (
	"html/template"
	"os"
)

type GitHubProfileData struct {
	Name string
}

type GitHubProfileTemplateData struct {
	Profile GitHubProfileData
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/gh-profile.html"))

	var data = GitHubProfileTemplateData{
		Profile: GitHubProfileData{
			Name: "Hugo Albrecht",
		},
	}

	err := tmpl.Execute(os.Stdout, data)

	if err != nil {
		return
	}
}