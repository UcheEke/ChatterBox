package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
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

	addr := ":8080"

	// Function
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// Start the web server
	fmt.Printf("Starting web server on localhost%s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
