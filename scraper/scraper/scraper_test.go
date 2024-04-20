package scraper

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func newUnstartedTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		file, _ := os.ReadFile("./testdata/full_document.html")

		w.Write(file)
	})

	return httptest.NewUnstartedServer(mux)
}

func newTestServer() *httptest.Server {
	srv := newUnstartedTestServer()
	srv.Start()
	return srv
}

func TestSingle(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	issueTime, _ := time.Parse(timeLayout, "2003-04-02")
	modifiedTime, _ := time.Parse(timeLayout, "2022-08-05")
	expectedArticle := Article{
		Title: "Michel Foucault",
		Author: []string{
			"Gutting, Gary",
			"Oksala, Johanna",
		},
		Issued:   issueTime,
		Modified: modifiedTime,
	}

	article, err := Single(ts.URL)
	if err != nil {
		t.Fatalf("Singe fetch failed: %v", err)
	}

	// Title
	if expectedArticle.Title != article.Title {
		t.Fatalf(
			"Wrong article title. Expect %s but got %s instead.",
			expectedArticle.Title,
			article.Title,
		)
	}

	// Author
	if len(article.Author) != len(expectedArticle.Author) {
		t.Fatalf(
			"Wrong author number. Expect %d but got %d instead.",
			len(expectedArticle.Author),
			len(article.Author),
		)
	}

	for i := range article.Author {
		if article.Author[i] != expectedArticle.Author[i] {
			t.Fatalf(
				"Wrong author. Expect %s but got %s instead",
				article.Author[i],
				expectedArticle.Author[i],
			)
		}
	}

	// Issued date
	if expectedArticle.Issued.Compare(article.Issued) != 0 {
		t.Fatalf("Wrong issued time. Expected %s but got %s instead",
			expectedArticle.Issued.Format(timeLayout),
			article.Issued.Format(timeLayout),
		)
	}

	// Modified date
	if expectedArticle.Modified.Compare(article.Modified) != 0 {
		t.Fatalf("Wrong modified time. Expected %s but got %s instead",
			expectedArticle.Modified.Format(timeLayout),
			article.Modified.Format(timeLayout),
		)
	}

	// Content
	if len(article.HTMLText) == 0 {
		t.Fatal("Empty HTML text")
	}

	if strings.Contains(article.HTMLText, `id="toc"`) {
		t.Fatalf("TOC is not filtered from HTML Text")
	}

	if strings.Contains(article.HTMLText, `id="academic-tools"`) {
		t.Fatalf("Academic tools is not filtered from HTML Text")
	}

}
