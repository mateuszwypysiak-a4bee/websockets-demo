package static

import (
	"html/template"
	"log"
	"net/http"
)

func HandleNative() func(w http.ResponseWriter, r *http.Request) {

	nativeTemplate, err := template.ParseFiles("templates/native.html")
	if err != nil {
		log.Fatal("Error parsing template templates/native.html\n", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		nativeTemplate.Execute(w, nil)

	}
}

func HandleStaticFiles() http.Handler {
	fileServer := http.FileServer(http.Dir("static/"))

	return http.StripPrefix("/static/", fileServer)
}
