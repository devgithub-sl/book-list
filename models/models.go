package models

// book data is represented as a struct
// this will normally represent a table in SQL if you think about it
type Book struct {
	Id     int
	Name   string
	Author string
	Isbn   string
}
