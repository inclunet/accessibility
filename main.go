package main

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/images"
	"github.com/inclunet/accessibility/pkg/report"
)

func GetPage(url string) (*goquery.Document, string, error) {
	Response, err := http.Get(url)

	if err != nil {
		return nil, "", err
	}

	defer Response.Body.Close()

	if Response.StatusCode != 200 {
		return nil, "", errors.New("URL not found")
	}

	Body, err := io.ReadAll(Response.Body)

	if err != nil {
		return nil, "", err
	}

	Html := string(Body)

	Response.Body = io.NopCloser(bytes.NewBuffer(Body))

	Document, err := goquery.NewDocumentFromReader(Response.Body)

	if err != nil {
		return nil, "", err
	}

	return Document, Html, nil
}

func CheckImages(document *goquery.Document, Checks *report.AccessibilityReport) {
	document.Find("img").Each(func(i int, s *goquery.Selection) {
		a, pass, description := images.Check(s)
		Checks.AddCheck(s, a, pass, description)
	})
}

func EvaluatePage(url string, lang string) {
	log.Printf("Starting page evaluation process for %s", url)
	Document, _, err := GetPage(url)

	if err != nil {
		log.Fatal(err)
	}

	title := Document.Find("title").Text()
	log.Printf("Evaluating page with title: %s", title)

	Checks := report.NewAccessibilityReport(url, title, lang)
	CheckImages(Document, &Checks)
	Checks.Save()
}

func main() {
	var url string
	var lang string
	flag.StringVar(&url, "url", "https://inclunet.com.br", "Url to check accessibility")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.Parse()
	EvaluatePage(url, lang)
}
