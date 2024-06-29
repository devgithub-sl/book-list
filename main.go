package main

import (
	"book/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// book data stored as map
// this map type acts like our temporary database for now
// technically we should use a database but for now we will use a map to store our data sets
var data = map[string][]models.Book{
	"Books": {models.Book{Id: 1, Name: "David Copperfield", Author: "Charles Dickens", Isbn: "978-0-141-34382-2"}},
}

// variable body to write multiple set of variables with data type assigned to it
var (
	// renders the book data from the map into the html via the template
	get_all_books = func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)
	}

	// add entered by our html form to be stored into our map
	add_new_book = func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.PostFormValue("name")
			author := r.PostFormValue("author")
			isbn := r.PostFormValue("isbn")

			book := models.Book{
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
)

func main() {
	fmt.Println("application is currently running on -> http://localhost:8080/")

	// routes
	http.HandleFunc("/", get_all_books)
	http.HandleFunc("/add-book", add_new_book)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
