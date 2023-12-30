package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
			{Title: "Se7ev", Director: "David Fincher"},
		},
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Print("Hello world received a request.")
		temp := template.Must(template.ParseFiles("index.html"))
		
		temp.Execute(w, films)
	}
	handle_add_film := func(w http.ResponseWriter, r *http.Request, ) {
		log.Print("handle_add_film received a request.")
		temp := template.Must(template.ParseFiles("add-film.html"))
		title := r.FormValue("title")
		director := r.FormValue("director")
		film := Film{Title: title, Director: director}
		to_send := map[string]Film{}
		to_send["Film"] = film
		films["Films"] = append(films["Films"], film)

		temp.Execute(w, to_send)
	}
	handle_delete := func(w http.ResponseWriter, r *http.Request, ) {
		log.Print("handle_delete received a request.")
		html_str := ""
		temp, _ := template.New("delete-film.html").Parse(html_str)
		
		temp.Execute(w, nil)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film", handle_add_film)
	http.HandleFunc("/delete-film", handle_delete)
	http.ListenAndServe(":8000", nil)
}
