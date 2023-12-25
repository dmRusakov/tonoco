package web

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
)

var _ Server = &server{}

func NewWebServer() (*server, error) {

	return &server{
		tmlPath: "./assets/templates/",
	}, nil
}

type server struct {
	tmlPath string
}

type Server interface {
	Render(w http.ResponseWriter, t string)
	Start(ctx context.Context) error
}

func (s server) Render(w http.ResponseWriter, pageTemplate string) {
	// page template
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s%s", s.tmlPath, pageTemplate))

	// static templates
	partials := []string{
		"base.layout",
		"head.partial",
		"header.partial",
		"footer.partial",
	}
	for _, x := range partials {
		templateSlice = append(templateSlice, fmt.Sprintf("%s%s.gohtml", s.tmlPath, x))
	}

	// parse templates
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// render page
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func (s server) Start(ctx context.Context) error {
	// static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// dynamic files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "test.page.gohtml")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
