package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func newUnstartedTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/entries/foucault", func(w http.ResponseWriter, r *http.Request) {
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
		EntryName: "foucault",
		Title:     "Michel Foucault",
		Author: []string{
			"Gutting, Gary",
			"Oksala, Johanna",
		},
		Issued:   issueTime,
		Modified: modifiedTime,
	}

	article, err := Single(ts.URL + "/entries/foucault")
	if err != nil {
		t.Fatalf("Singe fetch failed: %v", err)
	}

	// Entry name
	if expectedArticle.EntryName != article.EntryName {
		t.Fatalf(
			"Wrong article entry name. Expect %s but got %s instead.",
			expectedArticle.EntryName,
			article.EntryName,
		)
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
		t.Fatal("TOC is not filtered from HTML Text")
	}

	if strings.Contains(article.HTMLText, `id="academic-tools"`) {
		t.Fatal("Academic tools is not filtered from HTML Text")
	}

}

func compareTOCs(rootID string, t1 []TOCItem, t2 []TOCItem) error {
	if len(t1) != len(t2) {
		return fmt.Errorf(
			"Sub items length on %s mismatch. Expected %d but got %d",
			rootID,
			len(t1),
			len(t2),
		)
	}

	for i := range t1 {
		if t1[i].ID != t2[i].ID {
			return fmt.Errorf(
				"ID mismatch. Expected %s to equal %s.",
				t1[i].ID,
				t2[i].ID,
			)
		}

		if t1[i].Content != t2[i].Content {
			return fmt.Errorf(
				"Content mismatch. Expected \"%s\" to equal \"%s\".",
				t1[i].Content,
				t2[i].Content,
			)
		}

		if len(t1[i].SubItems) > 0 {
			err := compareTOCs(t1[i].ID, t1[i].SubItems, t2[i].SubItems)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func TestParseToc(t *testing.T) {
	file, _ := os.Open("./testdata/full_document.html")
	gq, _ := goquery.NewDocumentFromReader(file)

	testToc := []TOCItem{
		{
			ID:      "BiogSket",
			Content: "Biographical Sketch",
		},
		{
			ID:      "InteBack",
			Content: "Intellectual Background",
		},
		{
			ID:      "MajoWork",
			Content: "Major Works",
			SubItems: []TOCItem{
				{
					ID:      "HistMadnMedi",
					Content: "Histories of Madness and Medicine",
				},
				{
					ID:      "OrdeThin",
					Content: "The Order of Things",
					SubItems: []TOCItem{

						{
							ID:      "ClasRepr",
							Content: "Classical Representation",
						},
						{
							ID:      "KantCritClasRepr",
							Content: "Kant’s Critique of Classical Representation",
						},
						{
							ID:      "LangMan",
							Content: "Language and “Man”",
						},
						{
							ID:      "AnalFini",
							Content: "The Analytic of Finitude",
						},
					},
				},
				{
					ID:      "ArchGene",
					Content: "From Archaeology to Genealogy",
				},
				{
					ID:      "HistPris",
					Content: "History of the Prison",
				},
				{
					ID:      "HistModeSexu",
					Content: "History of Modern Sexuality",
				},
				{
					ID:      "SexAnciWorl",
					Content: "Sex in the Ancient World",
				},
			},
		},
		{
			ID:      "FoucAfteFouc",
			Content: "Foucault after Foucault",
		},
		{
			ID:      "Bib",
			Content: "Bibliography",
			SubItems: []TOCItem{
				{
					ID:      "PrimSour",
					Content: "Primary Sources",
				},
				{
					ID:      "SecoSour",
					Content: "Secondary Sources",
				},
			},
		},
		{
			ID:      "Oth",
			Content: "Other Internet Resources",
		},
		{
			ID:      "Rel",
			Content: "Related Entries",
		},
	}

	result := parseTOC(gq.Find("div[id='toc']").First())

	err := compareTOCs("root", testToc, result)
	if err != nil {
		t.Fatal(err)
	}
}
