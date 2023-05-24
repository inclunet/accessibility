package checker

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/accessibility"
	"github.com/inclunet/accessibility/pkg/report"
)

type CheckList struct {
	url        string
	ReportFile string
	Lang       string
}

type AccessibilityChecker struct {
	Domain     string
	Date       string
	Errors     int
	Pass       int
	Total      int
	TotalPages int
	CheckList  []CheckList
	Reports    []report.AccessibilityReport
}

func Check(s *goquery.Selection, accessibilityReport report.AccessibilityReport) report.AccessibilityReport {
	s.Each(func(i int, s *goquery.Selection) {
		elementName := goquery.NodeName(s)
		Html, _ := goquery.OuterHtml(s)
		A, Pass, Description, err := accessibility.NewElementCheck(s, accessibilityReport)

		if err == nil {
			accessibilityReport.AddCheck(elementName, A, Pass, Description, Html)
		}

		Check(s.Children(), accessibilityReport)
	})

	return accessibilityReport
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

func NewPageCheck(url string, reportFilename string, reportPath string, lang string) {
	log.Printf("Starting page evaluation process for %s", url)
	Document, _, err := GetPage(url)

	if err != nil {
		log.Fatal(err)
	}

	title := Document.Find("head title").Text()
	log.Printf("Evaluating page with title: %s", title)

	Report := Check(Document.Find("html"), report.NewAccessibilityReport(url, title, lang))

	log.Println("evaluation process is finished. Saving data on final report file...")
	err = Report.Save(reportPath, reportFilename)

	if err != nil {
		log.Fatal(err)
	} else {

		log.Printf("Evaluation report generated on %s file.", "a")
	}
}
