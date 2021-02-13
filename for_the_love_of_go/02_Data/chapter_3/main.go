package main

import "fmt"

// Book represents information about a book
type Book struct {
	Title           string
	Author          []string
	Series          string
	SeriesNumber    int
	Copies          int
	PriceCents      float64
	DiscountPercent float64
	FinalPrice      float64
	Edition         bool
}

// AddToCatalog takes a book and a catalog as parameters
// and adds the book to the catalog
func AddToCatalog(b Book, c []Book) []Book {
	c = append(c, b)
	fmt.Println(c)
	return c
}

func main() {

	var catalog []Book
	catalog = []Book{
		{Title: "The Shining",
			Author: []string{"A", "B"}},
		{Title: "IT",
			Author: []string{"C", "D"}},
	}
	b := Book{Title: "Third", Author: []string{"Mel Brooks", "Al Capone"}}

	catalog = AddToCatalog(b, catalog)

	fmt.Println(len(catalog))
	fmt.Println(catalog)
}
