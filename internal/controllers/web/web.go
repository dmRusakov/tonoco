package web

import (
	"fmt"
	"github.com/dmRusakov/tonoco/pkg/logrus"
	"html/template"
	"net/http"
)

var _ Server = &server{}

func NewWebServer(log *logrus.Logrus) (*server, error) {
	return &server{
		log:     log,
		tmlPath: "./assets/templates/",
	}, nil
}

type server struct {
	log     *logrus.Logrus
	tmlPath string
}

type Server interface {
	Render(w http.ResponseWriter, t string)
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
