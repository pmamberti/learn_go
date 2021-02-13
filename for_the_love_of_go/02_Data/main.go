package main

import "fmt"

// Book represents information about a book
type Book struct {
	Title           string
	Author          string
	Series          string
	SeriesNumber    int
	Copies          int
	PriceCents      float64
	DiscountPercent float64
	FinalPrice      float64
	Edition         bool
}

// NetPrice calculates the price of the book after the discount
func NetPrice(a Book) float64 {
	netDiscount := a.PriceCents * a.DiscountPercent / 100
	return (a.PriceCents - netDiscount) / 100
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

	fmt.Printf(`
	Author: %v
	Series: %v
	Title: %v
	Number in Series: %v
	
	Price: %v USD
	Discount: %v
	---------------
	Final Price: %.2f USD`, b.Author, b.Series, b.Title, b.SeriesNumber, b.PriceCents/100, b.DiscountPercent, b.FinalPrice)
}
