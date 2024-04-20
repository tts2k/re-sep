package scraper

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/microcosm-cc/bluemonday"
)

type TOCItem struct {
	ID       string
	Content  string
	SubItems []TOCItem
}

type Article struct {
	Title    string
	Issued   time.Time
	Modified time.Time
	HTMLText string
	Author   []string
	TOC      []TOCItem
}

var Collector = colly.NewCollector()

var timeLayout = "2006-01-02"
var htmlTrimRegex = regexp.MustCompile(`[\t\r\n]+`)
var tocContentRegex = regexp.MustCompile(`^[^a-zA-Z]+`)
var sanitizer = bluemonday.UGCPolicy()

func parseTOCRecur(root *goquery.Selection) []TOCItem {
	var toc []TOCItem

	root.Children().Each(func(_ int, sel *goquery.Selection) {
		tagName := goquery.NodeName(sel)

		if tagName != "li" {
			return
		}

		sel.Children().Each(func(_ int, subSel *goquery.Selection) {
			subTagName := goquery.NodeName(subSel)
			if subTagName == "a" {
				href, _ := subSel.Attr("href")

				// Skip Academic Tools
				if href == "#Aca" {
					return
				}

				// Trim space and new lines
				content := tocContentRegex.ReplaceAllString(strings.TrimSpace(subSel.Text()), "")
				content = htmlTrimRegex.ReplaceAllString(content, " ")

				toc = append(toc, TOCItem{
					ID:      href[1:],
					Content: content,
				})

				return
			}

			if subTagName == "ul" {
				// Find sub list
				currTocItem := &toc[len(toc)-1]
				currTocItem.SubItems = parseTOCRecur(subSel)
			}

		})

	})

	return toc
}

func parseTOC(root *goquery.Selection) []TOCItem {
	return parseTOCRecur(root.Find("ul").First())
}

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

	Collector.OnHTML("div[id='toc']", func(e *colly.HTMLElement) {
		article.TOC = parseTOC(e.DOM)
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
