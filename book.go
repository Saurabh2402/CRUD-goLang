package main

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Publisher struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

type Book struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Authors        []Author  `json:"authors"`
	PublishingInfo Publisher `json:"publisher"`
}
