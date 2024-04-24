package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// using colly to build this scraper, it can be used to build crawlers or spiders
// Colly is an open-source library that provides a clean interface based on callbacks to
// write a scraper, crawler or spider. It comes with an advanced Go web scraping API that allows you to download an HTML page,
//  automatically parse its content, select HTML elements from the DOM and retrieve data from them.

func main() {

	collector := colly.NewCollector()
	collector.Visit("https://gohugo.io/content-management/toc/")

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL)
	})

	collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visiting: ", r.Request.URL)
	})

	collector.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Println("%v", e.Attr("href"))
	})

	collector.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

}
