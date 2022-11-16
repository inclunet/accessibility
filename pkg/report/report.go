package report

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/summary"
)

type AccessibilityReport struct {
	URL    string
	Title  string
	Lang   string
	Checks []summary.AccessibilityCheck
}

func (r *AccessibilityReport) AddCheck(s *goquery.Selection, a int, pass bool, description string) {
	element := goquery.NodeName(s)
	html, _ := goquery.OuterHtml(s)
	r.Checks = append(r.Checks, summary.AccessibilityCheck{Element: element, A: a, Pass: pass, Description: description, Html: html})
}

func (r *AccessibilityReport) Save() {
	for _, entry := range summary.Generate(r.Checks) {
		log.Printf("%d %s tested with %d errors and %d asserts", entry.Total, entry.Element, entry.Errors, entry.Pass)
	}
}

func NewAccessibilityReport(url string, title string, lang string) AccessibilityReport {
	return AccessibilityReport{URL: url, Title: title, Lang: lang}
}
