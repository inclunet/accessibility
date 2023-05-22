package accessibility

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
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

func Check(s *goquery.Selection, accessibilityReport *report.AccessibilityReport) {
	s.Each(func(i int, s *goquery.Selection) {
		elementName := goquery.NodeName(s)
		Html, _ := goquery.OuterHtml(s)
		A, Pass, Description, err := NewCheck(s, accessibilityReport)
		if err == nil {
			accessibilityReport.AddCheck(elementName, A, Pass, Description, Html)
		}
		Check(s.Children(), accessibilityReport)
	})
}

func EvaluatePage(url string, reportFilename string, lang string) {
	log.Printf("Starting page evaluation process for %s", url)
	Document, _, err := GetPage(url)

	if err != nil {
		log.Fatal(err)
	}

	title := Document.Find("head title").Text()
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

func DeepCheck(s *goquery.Selection, accessibilityReport *report.AccessibilityReport) (int, bool, string, error) {
	A, Pass, Description, err := NewCheck(s, accessibilityReport)
	if err != nil {
		s.Each(func(i int, s *goquery.Selection) {
			A, Pass, Description, err = DeepCheck(s.Children(), accessibilityReport)
		})
	}
	return A, Pass, Description, err
}

func NewCheck(s *goquery.Selection, accessibilityReport *report.AccessibilityReport) (int, bool, string, error) {
	accessibilityCheckList := map[string]func(*goquery.Selection, *report.AccessibilityReport) Accessibility{
		"a":      NewLinkCheck,
		"button": NewButtonCheck,
		"h1":     NewHeaderCheck,
		"h2":     NewHeaderCheck,
		"h3":     NewHeaderCheck,
		"h4":     NewHeaderCheck,
		"h5":     NewHeaderCheck,
		"h6":     NewHeaderCheck,
		"input":  NewInputCheck,
		"img":    NewImageCheck,
	}
	elementName := goquery.NodeName(s)
	if elementInterface, ok := accessibilityCheckList[elementName]; ok {
		accessibilityInterface := elementInterface(s, accessibilityReport)
		a, pass, description := accessibilityInterface.Check()
		return a, pass, description, nil
	}
	return 1, false, "", errors.New("no defined evaluator function for this html element")
}
