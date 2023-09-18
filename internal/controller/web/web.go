package web

import (
	"fmt"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
	"html/template"
	"net/http"
)

var _ Server = &server{}

func NewWebServer(log *logrus.Logrus) (*server, error) {
	return &server{
		log: log,
	}, nil
}

type server struct {
	log *logrus.Logrus
}

type Server interface {
	Render(w http.ResponseWriter, t string)
}

func (s server) Render(w http.ResponseWriter, t string) {

	partials := []string{
		"./assets/templates/base.layout.gohtml",
		"./assets/templates/header.partial.gohtml",
		"./assets/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./assets/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
