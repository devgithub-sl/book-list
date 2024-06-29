package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// book data is represented as a struct
// this will normally represent a table in SQL if you think about it
type Book struct {
	Id     int
	Name   string
	Author string
	Isbn   string
}

func main() {
	fmt.Println("herro")

	// this map type acts like our temporary database for now
	// book data stored as map
	// technically we should use a database but for now we will use a map to store our data sets
	data := map[string][]Book{
		"Books": {
			Book{
				Id: 1, Name: "David Copperfield", Author: "Charles Dickens", Isbn: "978-0-141-34382-2",
			},
		},
	}

	// renders the book data from the map into the html via the template
	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)
	}

	// add entered by our html form to be stored into our map
	addBookHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.PostFormValue("name")
			author := r.PostFormValue("author")
			isbn := r.PostFormValue("isbn")

			book := Book{
				Id:     len(data["Books"]) + 1,
				Name:   name,
				Author: author,
				Isbn:   isbn,
			}
			data["Books"] = append(data["Books"], book)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"hx-redirect": "/",
			})
		}
		templ := template.Must(template.ParseFiles("add-book.html"))
		templ.Execute(w, nil)
	}

	// routes
	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-book", addBookHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
