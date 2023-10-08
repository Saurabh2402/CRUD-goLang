package main

type Author struct {
	FirstName string
	LastName  string
}

type Publisher struct {
	Name string
	Year int
}

type Book struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Authors        []Author
	PublishingInfo Publisher
}
