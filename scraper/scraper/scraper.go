package scraper

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/microcosm-cc/bluemonday"
)

type TOCItem struct {
	ID       string    `json:"id"`
	Label    string    `json:"label"`
	SubItems []TOCItem `json:"subItems"`
}

type Article struct {
	EntryName string    `json:"entryName"`
	Title     string    `json:"title"`
	Issued    time.Time `json:"issued"`
	Modified  time.Time `json:"modified"`
	HTMLText  []byte    `json:"htmlText"`
	Author    []string  `json:"author"`
	TOC       []TOCItem `json:"toc"`
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
					ID:    href[1:],
					Label: content,
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

func addCSSTemplateTags(dom *goquery.Selection) {
	dom.Find("h1").AddClass("{{h1}}")
	dom.Find("h2").AddClass("{{h2}}")
	dom.Find("h3").AddClass("{{h3}}")
	dom.Find("h4").AddClass("{{h4}}")
	dom.Find("p, ul, em").AddClass("{{text}}")
}

func Single(url string) (Article, error) {
	article := Article{}

	article.EntryName = path.Base(url)

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

		addCSSTemplateTags(dom)

		HTMLText, err := dom.Html()
		if err != nil {
			fmt.Printf("Error extracting content: %v\n", err)
			return
		}

		// Trim whiltespaces
		HTMLText = htmlTrimRegex.ReplaceAllString(strings.TrimSpace(HTMLText), "\n")

		// Sanitize
		HTMLText = sanitizer.Sanitize(HTMLText)

		// Compress
		var b bytes.Buffer

		gz := gzip.NewWriter(&b)
		gz.Write([]byte(HTMLText))
		gz.Close()

		article.HTMLText = b.Bytes()
	})

	err := Collector.Visit(url)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}
