package main

import (
	"html/template"
	"log"
	"net/http"
	"photo-site/internal/storage"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	storage.Init("./photo-site.db")

	token, err := storage.CreateEvent("Свадьба Ивановых", "wedding-ivan")
	if err != nil {
		log.Println("Ошибка создания мероприятия:", err)
	} else {
		log.Println("Ссылка на мероприятие: http://localhost:7000/event/" + token)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/about/", aboutPage)
	http.HandleFunc("/event/", eventPage)

	log.Println("Server Up, http:/localhost:7000")

	log.Fatal(http.ListenAndServe(":7000", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "template", http.StatusInternalServerError)
	}
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Error(w, "template", http.StatusInternalServerError)
	}
}

func eventPage(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Path[len("/event/"):]
	if token == "" {
		http.NotFound(w, r)
		return
	}

	event, err := storage.GetEventByToken(token)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	err = templates.ExecuteTemplate(w, "event.html", event)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		log.Println("eventPage error:", err)
	}

}
