package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func main() {
	// TODO: Implementare il concurrent web scraper
	fmt.Println("Concurrent Web Scraper")
}

func extractTitleAndLinks(r io.Reader) (string, int) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", 0
	}

	var title string
	linkCount := 0

	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			linkCount++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)

	return title, linkCount
}
