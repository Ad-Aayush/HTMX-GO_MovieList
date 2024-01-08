package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type Film struct {
	ID       int
	Title    string
	Director string
}

var (
	films  = make(map[string]map[int]Film)
	nextID = 3
)

func main() {
	// Initialize some films
	if films["Films"] == nil {
		films["Films"] = make(map[int]Film)
	}
	films["Films"][0] = Film{ID: 0, Title: "The Godfather", Director: "Francis Ford Coppola"}
	films["Films"][1] = Film{ID: 1, Title: "Pulp Fiction", Director: "Quentin Tarantino"}
	films["Films"][2] = Film{ID: 2, Title: "Se7en", Director: "David Fincher"}

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film", handleAddFilm)
	http.HandleFunc("/edit-film", handleEdit)
	http.HandleFunc("/delete-film", handleDeleteFilm)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Home Page received a request.")
	temp := template.Must(template.ParseFiles("index.html"))
	temp.Execute(w, films)
}

func handleAddFilm(w http.ResponseWriter, r *http.Request) {
	log.Print("Add Film received a request.")
	title := r.FormValue("title")
	director := r.FormValue("director")
	film := Film{ID: nextID, Title: title, Director: director}
	films["Films"][nextID] = film
	nextID++

	temp := template.Must(template.ParseFiles("add-film.html"))
	temp.Execute(w, film)
}

func handleDeleteFilm(w http.ResponseWriter, r *http.Request) {
	// log.Print("Delete Film received a request.")
	idStr := r.FormValue("filmId")
	id, _ := strconv.Atoi(idStr)
	// log.Print("Size:", len(films["Films"]))
	delete(films["Films"], id)
	// log.Print("DeletedId: ", id)

	temp := template.Must(template.ParseFiles("add-film.html"))
	// Keys slice
	keySlice := []int{}

	for key := range films["Films"] {
		keySlice = append(keySlice, key)
		// log.Print("Key: ", key, ", Film: ", film)
	}

	sort.Ints(keySlice)
	log.Printf("KeySlice: %v", keySlice)
	for _, key := range keySlice {

		temp.Execute(w, films["Films"][key])
	}
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	log.Print("Edit Film received a request.")
	idStr := r.FormValue("filmId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == http.MethodPost {
		film := films["Films"][id]
		temp := template.Must(template.ParseFiles("edit-film.html"))
		log.Print("POST FILM", film)
		temp.Execute(w, film)
	} else if r.Method == http.MethodPut {
		log.Print("Edit Film PUT received a request.")

		log.Print("ID: ", idStr)
		title := r.FormValue("title")
		director := r.FormValue("director")
		film := Film{ID: id, Title: title, Director: director}
		films["Films"][id] = film
		log.Print("Film: ", film, idStr, id)
		temp := template.Must(template.ParseFiles("add-film.html"))
		temp.ExecuteTemplate(w, "edit-film", film)
	} else if r.Method == http.MethodGet {
		log.Print("Edit Film GET received a request.")

		film := films["Films"][id]
		temp := template.Must(template.ParseFiles("add-film.html"))
		temp.ExecuteTemplate(w, "edit-film", film)
	}
}
