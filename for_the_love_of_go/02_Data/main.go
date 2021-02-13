package main

import "fmt"

func main() {

	// Book represents information about a book
	type Book struct {
		Title        string
		Author       string
		Series       string
		SeriesNumber int
	}

	var b = Book{
		Title:        "Fundamentals",
		Author:       "John Arundel",
		Series:       "For the Love of Go",
		SeriesNumber: 1,
	}

	fmt.Println(b)
}
