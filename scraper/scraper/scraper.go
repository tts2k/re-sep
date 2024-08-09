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

	"re-sep-scraper/config"
	"re-sep-scraper/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/microcosm-cc/bluemonday"
)

type TOCItem struct {
	ID       string    `json:"id"`
	Label    string    `json:"label"`
	SubItems []TOCItem `json:"subItems"`
}

type HTMLText = string

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
	dom.Find("p, ul, em, blockquote").AddClass("{{text}}")
}

func SingleContentOnly(url string) (string, error) {
	var result string
	collector := colly.NewCollector()

	collector.OnHTML("div[id='aueditable']", func(e *colly.HTMLElement) {
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

		result = HTMLText
	})

	err := collector.Visit(url)

	if err != nil {
		return result, err
	}

	return result, nil

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

		HTMLText, err := dom.Html()
		if err != nil {
			fmt.Printf("Error extracting content: %v\n", err)
			return
		}

		// Sanitize
		HTMLText = sanitizer.Sanitize(HTMLText)

		newDoc, err := goquery.NewDocumentFromReader(strings.NewReader(HTMLText))
		if err != nil {
			fmt.Printf("Error parsing sanitized content: %v\n", err)
			return
		}

		addCSSTemplateTags(newDoc.Selection)

		HTMLText, err = newDoc.Html()
		if err != nil {
			fmt.Printf("Error parsing sanitized content: %v\n", err)
			return
		}

		// Compress
		var b bytes.Buffer

		gz := gzip.NewWriter(&b)
		_, _ = gz.Write([]byte(HTMLText))
		gz.Close()

		article.HTMLText = b.Bytes()
	})

	err := collector.Visit(url)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}

func spawnWorkers(wg *sync.WaitGroup, jobs <-chan string, results chan<- Article) {
	for i := 0; i < config.WorkerCount; i++ {
		wg.Add(1)
		go func() {
			for j := range jobs {
				utils.Debugln("Working on:", j)
				res, err := Single("https://plato.stanford.edu/" + j)
				if err != nil {
					utils.Debugln(err)
					continue
				}

				results <- res
				time.Sleep(time.Duration(config.Sleep) * time.Millisecond)
			}
			wg.Done()
		}()
	}
}

func All() (*sync.WaitGroup, chan Article, error) {
	// Just scrape the TOC. No need for bfs/dfs scraping!
	url := "https://plato.stanford.edu/contents.html"
	collector := colly.NewCollector()
	jobs := make(chan string, 100)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		var count int
		visited := make(map[string]bool)
		collector.OnHTML(`a[href^="entries"]`, func(e *colly.HTMLElement) {
			// Hard-coded arbitrary limit for development
			// I don't wanna scrape 1800 articles to test
			if count >= 20 {
				return
			}

			href := e.Attr("href")
			// Skip visited sites
			if visited[href] {
				utils.Debugf("Skipping visited entry: %s", href)
				return
			}

			if href != "" {
				jobs <- href
				visited[href] = true
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

	results := make(chan Article, 100)
	spawnWorkers(wg, jobs, results)

	return wg, results, nil
}
