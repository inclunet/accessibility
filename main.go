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
	"github.com/inclunet/accessibility/pkg/summary"
)

var AccessibilityChecks []summary.AccessibilityCheck
var Html string

func AccessibilityCheckResult(s *goquery.Selection, a int, pass bool, description string) {
	element := goquery.NodeName(s)
	html, _ := goquery.OuterHtml(s)
	AccessibilityChecks = append(AccessibilityChecks, summary.AccessibilityCheck{Element: element, A: a, Pass: pass, Description: description, Html: html})
}

func CheckImageAccessibility(i int, s *goquery.Selection) {
	a, pass, description := images.Check(s)
	AccessibilityCheckResult(s, a, pass, description)
}

func GetPage(url string) (*http.Response, error) {
	Response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer Response.Body.Close()

	if Response.StatusCode != 200 {
		return nil, errors.New("URL not found")
	}

	Body, err := io.ReadAll(Response.Body)

	if err != nil {
		return nil, err
	}

	Html = string(Body)

	Response.Body = io.NopCloser(bytes.NewBuffer(Body))

	return Response, nil
}

func SaveResults() {
	for _, entry := range summary.Generate(AccessibilityChecks) {
		log.Printf("%d %s tested with %d errors and %d asserts", entry.Total, entry.Element, entry.Errors, entry.Pass)
	}
}

func EvaluatePage(url string, lang string) {
	log.Printf("Starting page evaluation process for %s", url)

	Response, err := GetPage(url)

	if err != nil {
		log.Fatal(err)
	}

	document, err := goquery.NewDocumentFromReader(Response.Body)

	if err != nil {
		log.Fatal(err)
	}

	title := document.Find("title").Text()

	log.Printf("Evaluating page with title: %s", title)

	document.Find("img").Each(CheckImageAccessibility)

	SaveResults()
}

func main() {
	var url string
	var lang string
	flag.StringVar(&url, "url", "https://inclunet.com.br", "Url to check accessibility")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.Parse()
	EvaluatePage(url, lang)
}
