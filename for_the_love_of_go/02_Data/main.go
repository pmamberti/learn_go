package main

import "fmt"

// Book represents information about a book
type Book struct {
	Title           string
	Author          string
	Series          string
	SeriesNumber    int
	Copies          int
	PriceCents      int
	DiscountPercent int
	FinalPrice      int
	Edition         bool
}

// NetPrice calculates the price of the book after the discount
func NetPrice(a Book) int {
	netDiscount := a.PriceCents * a.DiscountPercent / 100
	return a.PriceCents - netDiscount
}

func main() {

	b := Book{
		Title:           "Fundamentals",
		Author:          "John Arundel",
		Series:          "For the Love of Go",
		SeriesNumber:    1,
		Copies:          7,
		PriceCents:      499,
		DiscountPercent: 10,
	}

	b.FinalPrice = NetPrice(b)

	fmt.Printf("Author: %v\nSeries: %v\nTitle: %v\nNumber in Series: %v\nPrice: %v USD\nDiscount: %v%\nFinal Price: %v\n", b.Author, b.Series, b.Title, b.SeriesNumber, b.PriceCents/100, b.DiscountPercent, b.FinalPrice)
	fmt.Println(b.Edition)
}
