package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Name   string
	Price  float64
	Rating float64
}

var Products []Product

func compareRating(i int, j int) bool {
	if Products[i].Rating == Products[j].Rating {
		return Products[i].Price < Products[j].Price
	}
	return Products[i].Rating > Products[j].Rating
}

func main() {

	c := colly.NewCollector((colly.AllowedDomains("www.amazon.in")))

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Link of the page:", r.URL, "\n \n")
	})

	c.OnHTML("div.s-main-slot.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(_ int, h *colly.HTMLElement) {
			name := h.ChildText("span.a-color-base.a-text-normal")
			stars := h.ChildText("span.a-icon-alt")
			price := h.ChildText("span.a-price-whole")

			var temp Product

			if name != "" {
				s := strings.Split(stars, " ")
				price_Float, _ := strconv.ParseFloat(price, 64)
				rating_Float, _ := strconv.ParseFloat(s[0], 64)

				temp.Name = name
				temp.Rating = rating_Float
				temp.Price = price_Float
				Products = append(Products, temp)
			}
		})
	})
	var keyword string = "shoes"
	c.Visit("https://www.amazon.in/s?k=" + keyword)

	sort.Slice(Products, compareRating)

	fmt.Print("Top 5 Products for the given keyword: \n")
	fmt.Println(Products)
	if len(Products) == 0 {
		fmt.Println("Unable to access website :(")
		return
	}

	for i := int64(0); float64(i) < math.Min(float64(len(Products)), float64(5)); i++ {

		fmt.Println("Ranking: #", i+1)
		fmt.Println("Name:", Products[i].Name)
		fmt.Println("Rating:", Products[i].Rating, "/n")
	}
}
