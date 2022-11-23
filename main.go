package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/accessibility"
	"github.com/inclunet/accessibility/pkg/report"
)

func a(s *goquery.Selection, Report *report.AccessibilityReport) (int, bool, string) {
	fmt.Println(goquery.NodeName(s))
	return 1, false, ""
}

func CheckList(s *goquery.Selection, Report *report.AccessibilityReport) {
	fnList := map[string]func(*goquery.Selection) (int, bool, string){
		"img": accessibility.NewImageCheck,
	}
	Element := goquery.NodeName(s)
	if fn, ok := fnList[Element]; ok {
		Html, _ := goquery.OuterHtml(s)
		A, Pass, Description := fn(s)
		Report.AddCheck(Element, A, Pass, Description, Html)
	}
}

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

func Check(s *goquery.Selection, Report *report.AccessibilityReport) {
	s.Each(func(i int, s *goquery.Selection) {
		CheckList(s, Report)
		Check(s.Children(), Report)
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

	Report := report.NewAccessibilityReport(url, title, lang)
	Check(Document.Find("html"), &Report)
	Report.Save()
}

func main() {
	var url string
	var lang string
	flag.StringVar(&url, "url", "https://inclunet.com.br", "Url to check accessibility")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.Parse()
	EvaluatePage(url, lang)
}
