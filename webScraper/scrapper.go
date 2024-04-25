package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// using colly to build this scraper, it can be used to build crawlers or spiders
// Colly is an open-source library that provides a clean interface based on callbacks to
// write a scraper, crawler or spider. It comes with an advanced Go web scraping API that allows you to download an HTML page,
//  automatically parse its content, select HTML elements from the DOM and retrieve data from them.

func main() {

	fmt.Println("Scraping web...")
	c := colly.NewCollector(
		colly.AllowedDomains("https://en.wikipedia.org/wiki/Civil_War_(film)"),
	)

	err := c.Visit("https://en.wikipedia.org/wiki/Civil_War_(film)")

	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return
	}

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visiting: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Printf("%v", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.Wait()

}
