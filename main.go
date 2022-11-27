package main

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/accessibility"
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

func Check(s *goquery.Selection, Report *report.AccessibilityReport) {
	s.Each(func(i int, s *goquery.Selection) {
		elementName := goquery.NodeName(s)
		Html, _ := goquery.OuterHtml(s)
		A, Pass, Description, err := accessibility.NewCheck(s)
		if err == nil {
			Report.AddCheck(elementName, A, Pass, Description, Html)
		}
		Check(s.Children(), Report)
	})
}

func EvaluatePage(url string, reportFilename string, lang string) {
	log.Printf("Starting page evaluation process for %s", url)
	Document, _, err := GetPage(url)

	if err != nil {
		log.Fatal(err)
	}

	title := Document.Find("title").Text()
	log.Printf("Evaluating page with title: %s", title)

	Report := report.NewAccessibilityReport(url, title, lang)
	Check(Document.Find("html"), &Report)

	log.Println("evaluation process is finished. Saving data on final report file...")
	err = Report.Save(reportFilename)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Evaluation report generated on %s file.", reportFilename)
	}
}

func main() {
	var url string
	var report string
	var lang string
	flag.StringVar(&url, "url", "https://inclunet.com.br", "Url to check accessibility")
	flag.StringVar(&report, "report", "report.html", "Output filename to generate final html report")
	flag.StringVar(&lang, "lang", "pt-BR", "Content language of page")
	flag.Parse()
	EvaluatePage(url, report, lang)
}
