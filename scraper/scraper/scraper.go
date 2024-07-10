package scraper

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"path"
	"regexp"
	"strings"
	"sync"
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
	collector := colly.NewCollector()

	article := Article{}

	article.EntryName = path.Base(url)

	collector.OnHTML(`meta[name="DC.title"]`, func(e *colly.HTMLElement) {
		article.Title = e.Attr("content")
	})

	collector.OnHTML(`meta[name="DC.creator"]`, func(e *colly.HTMLElement) {
		article.Author = append(article.Author, e.Attr("content"))
	})

	collector.OnHTML(`meta[name="DCTERMS.issued"]`, func(e *colly.HTMLElement) {
		issuedTime, err := time.Parse(timeLayout, e.Attr("content"))
		if err != nil {
			fmt.Printf("Error parsing time: %s\n", e.Attr("content"))
			return
		}

		article.Issued = issuedTime
	})

	collector.OnHTML(`meta[name="DCTERMS.modified"]`, func(e *colly.HTMLElement) {
		modifiedTime, err := time.Parse(timeLayout, e.Attr("content"))
		if err != nil {
			fmt.Printf("Error parsing time: %s\n", e.Attr("content"))
			return
		}

		article.Modified = modifiedTime
	})

	collector.OnHTML("div[id='toc']", func(e *colly.HTMLElement) {
		article.TOC = parseTOC(e.DOM)
	})

	collector.OnHTML("div[id='aueditable']", func(e *colly.HTMLElement) {
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

	err := collector.Visit(url)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}

func createSafePrintln() func(a ...any) {
	printLock := &sync.Mutex{}

	return func(a ...any) {
		printLock.Lock()
		fmt.Println(a...)
		printLock.Unlock()
	}
}

func spawnWorkers(wCount int, wg *sync.WaitGroup, jobs <-chan string, results chan<- Article) {
	safePrintln := createSafePrintln()

	for i := 0; i < wCount; i++ {
		wg.Add(1)
		go func() {
			for j := range jobs {
				safePrintln("Working on:", j)
				res, err := Single("https://plato.stanford.edu/" + j)
				if err != nil {
					safePrintln(err)
					continue
				}

				results <- res
			}
			wg.Done()
		}()
	}
}

func All() (*sync.WaitGroup, chan Article, error) {
	url := "https://plato.stanford.edu/contents.html"
	collector := colly.NewCollector()
	jobs := make(chan string, 100)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		var count int
		collector.OnHTML(`a[href^="entries"]`, func(e *colly.HTMLElement) {
			if count >= 20 {
				return
			}
			href := e.Attr("href")
			if href != "" {
				jobs <- href
				count++
			}
		})

		err := collector.Visit(url)

		if err != nil {
			panic(err)
		}
		close(jobs)
		wg.Done()
	}()

	workerCount := 5
	results := make(chan Article, 100)
	spawnWorkers(workerCount, wg, jobs, results)

	return wg, results, nil
}
