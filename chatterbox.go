package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
	"github.com/codegangsta/negroni"
	gmux "github.com/gorilla/mux"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the http request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	mux := gmux.NewRouter()
	addr := "localhost:8080"
	r := newRoom()

	mux.Handle("/", &templateHandler{filename: "chat.html"}).Methods("GET")
	mux.Handle("/room", r).Methods("GET")

	// Start the room running as a concurrent process
	go r.run()

	// Handle the static files (*.js,*.css, images etc)
	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Create a negroni Classic middleware handler
	n := negroni.Classic()

	n.UseHandler(mux)
	// Start the listening on the user specified port
	n.Run(addr)

}