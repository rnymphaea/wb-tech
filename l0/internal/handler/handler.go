package handler

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed templates/*.html
var templateFS embed.FS

var templates *template.Template

func init() {
	loadTemplates()
}

func loadTemplates() {
	subFS, err := fs.Sub(templateFS, "templates")
	if err != nil {
		log.Fatalf("Error accessing templates subdirectory: %v", err)
	}

	tmpl, err := template.ParseFS(subFS, "*.html")
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}
	templates = tmpl
  if err := fs.WalkDir(templateFS, ".", func(path string, d fs.DirEntry, err error) error {
        log.Printf("Found file in FS: %s", path)
        return nil
    }); err != nil {
        log.Printf("Error walking FS: %v", err)
    }
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "Title": "Order Lookup",
        "Error": r.URL.Query().Get("error"),
    }

    err := templates.ExecuteTemplate(w, "index.html", data)
    if err != nil {
        log.Printf("Error rendering home page: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}
