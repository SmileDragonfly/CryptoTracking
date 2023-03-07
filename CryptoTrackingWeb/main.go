package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/WelcomeTemplate.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "WelcomeTemplate.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start web server, và sét cổng 8080
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
