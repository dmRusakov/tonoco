package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// start web router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		App.Router.Web.Render(w, "test.page.gohtml")
	})

	fmt.Printf("Starting front end service on port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Print(App)
}
