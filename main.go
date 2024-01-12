package main

import (
	"context"
	"html/template"
	"log"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
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

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		log.Print("Home Page received a request.")
		// temp := template.Must(template.ParseFiles("index.html"))
		keySlice := []int{}

		for key := range films["Films"] {
			keySlice = append(keySlice, key)
			// log.Print("Key: ", key, ", Film: ", film)
		}

		sort.Ints(keySlice)
		component := index(films["Films"], keySlice)
		return component.Render(context.Background(), c.Response().Writer)
	})
	// http.HandleFunc("/add-film", handleAddFilm)
	e.POST("/add-film", func(c echo.Context) error {
		log.Print("Add Film received a request.")
		title := c.FormValue("title")
		director := c.FormValue("director")
		film := Film{ID: nextID, Title: title, Director: director}
		films["Films"][nextID] = film
		nextID++

		// temp := template.Must(template.ParseFiles("add-film.html"))
		// return temp.Execute(c.Response().Writer, film)
		component := add_film(film)
		return component.Render(context.Background(), c.Response().Writer)
	})
	// http.HandleFunc("/edit-film", handleEdit)
	e.POST("/edit-film", func(c echo.Context) error {
		log.Print("Edit Film received a request.")
		idStr := c.FormValue("filmId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}

		film := films["Films"][id]
		log.Print("POST FILM", film)

		// temp := template.Must(template.ParseFiles("edit-film.html"))
		// return temp.Execute(c.Response().Writer, film)
		component := edit(film)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.PUT("/edit-film", func(c echo.Context) error {
		log.Print("Edit Film received a request.")
		idStr := c.FormValue("filmId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(c, err)
		}

		log.Print("Edit Film PUT received a request.")

		log.Print("ID: ", idStr)
		title := c.FormValue("title")
		director := c.FormValue("director")
		film := Film{ID: id, Title: title, Director: director}
		films["Films"][id] = film
		log.Print("Film: ", film, idStr, id)
		// temp := template.Must(template.ParseFiles("add-film.html"))
		// return temp.ExecuteTemplate(c.Response().Writer, "edit-film", film)
		component := edit_rep(film)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.PUT("/can-film", func(c echo.Context) error {
		log.Print("Edit Film GET received a request.")

		idStr := c.FormValue("filmId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(c, err)
		}
		film := films["Films"][id]
		// temp := template.Must(template.ParseFiles("add-film.html"))
		// return temp.ExecuteTemplate(c.Response().Writer, "edit-film", film)
		component := edit_rep(film)
		return component.Render(context.Background(), c.Response().Writer)
	})

	// http.HandleFunc("/delete-film", handleDeleteFilm)
	e.POST("/delete-film", func(c echo.Context) error {
		idStr := c.FormValue("filmId")
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

			temp.Execute(c.Response().Writer, films["Films"][key])
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":8000"))
}
