package webui

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// TemplatesDir is the directory where the HTML templates are stored.
var TemplatesDir = "templates" // Adjust the path according to your project structure

// RenderIndex handles rendering of the index page.
func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmplPath := path.Join(TemplatesDir, "index.gohtml")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Bitcoin Explorer",
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
