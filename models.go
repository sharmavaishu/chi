package main



type Book struct {
	ID         string `json : "id"`
	Title        string  `json: "title"`
	Author      string  `json: "author"`
	ReleaseDate string  `json: "ReleaseDate"`
}

var books = []*Book{
	{
	ID:               "1",
	Title:            "HarryPorter",
	Author:           "Jk rowling",
	ReleaseDate:    "15/08/1989",
	},
}

func listBooks() []*Book{
     return books
}

