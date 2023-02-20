package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Welcome struct {
	Time string
}

func main() {
	welcome := Welcome{time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("html/template.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}
