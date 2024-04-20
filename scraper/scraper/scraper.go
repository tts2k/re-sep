package scraper

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/microcosm-cc/bluemonday"
)

type Article = struct {
	Title    string
	Issued   time.Time
	Modified time.Time
	HTMLText string
	Author   []string
}

var Collector = colly.NewCollector()

var timeLayout = "2006-01-02"
var htmlTrimRegex = regexp.MustCompile(`[\t\r\n]+`)
var sanitizer = bluemonday.UGCPolicy()

func Single(url string) (Article, error) {
	article := Article{}

	Collector.OnHTML(`meta[name="DC.title"]`, func(e *colly.HTMLElement) {
		article.Title = e.Attr("content")
	})

	Collector.OnHTML(`meta[name="DC.creator"]`, func(e *colly.HTMLElement) {
		article.Author = append(article.Author, e.Attr("content"))
	})

	Collector.OnHTML(`meta[name="DCTERMS.issued"]`, func(e *colly.HTMLElement) {
		issuedTime, err := time.Parse(timeLayout, e.Attr("content"))
		if err != nil {
			fmt.Printf("Error parsing time: %s\n", e.Attr("content"))
			return
		}

		article.Issued = issuedTime
	})

	Collector.OnHTML(`meta[name="DCTERMS.modified"]`, func(e *colly.HTMLElement) {
		modifiedTime, err := time.Parse(timeLayout, e.Attr("content"))
		if err != nil {
			fmt.Printf("Error parsing time: %s\n", e.Attr("content"))
			return
		}

		article.Modified = modifiedTime
	})

	Collector.OnHTML("div[id='aueditable']", func(e *colly.HTMLElement) {
		dom := e.DOM
		dom.Find("#toc").Remove()
		dom.Find("#academic-tools").Remove()

		HTMLText, err := dom.Html()
		if err != nil {
			fmt.Printf("Error extracting content: %v\n", err)
			return
		}

		// Trim whiltespaces
		HTMLText = htmlTrimRegex.ReplaceAllString(strings.TrimSpace(HTMLText), "\n")

		// Sanitize
		HTMLText = sanitizer.Sanitize(HTMLText)

		article.HTMLText = HTMLText
	})

	err := Collector.Visit(url)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}
