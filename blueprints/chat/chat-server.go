package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	port := 8080
	args := os.Args[1:]
	if len(args) > 0 {
		port, _ = strconv.Atoi(args[0])
	}

	http.Handle("/", &templateHandler{filename: "chat.html"})

	address := ":" + strconv.Itoa(port)
	log.Println("Listening on", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
