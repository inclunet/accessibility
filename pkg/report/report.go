package report

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

type AccessibilityCheck struct {
	Element     string
	A           int
	Pass        bool
	Description string
	Line        int
	Html        string
}

type AccessibilityReport struct {
	URL     string
	Title   string
	Lang    string
	Pass    int
	Errors  int
	Total   int
	Summary map[string]*AccessibilitySummary
	Checks  []AccessibilityCheck
}

func (r *AccessibilityReport) AddCheck(s *goquery.Selection, a int, pass bool, description string) {
	element := goquery.NodeName(s)
	html, _ := goquery.OuterHtml(s)
	r.Checks = append(r.Checks, AccessibilityCheck{Element: element, A: a, Pass: pass, Description: description, Html: html})
}

func (r *AccessibilityReport) GenerateSummary() {
	for _, check := range r.Checks {
		_, ok := r.Summary[check.Element]
		if !ok {
			r.Summary[check.Element] = NewSummary()
		}
		if check.Pass {
			r.Summary[check.Element].AddPass()
		} else {
			r.Summary[check.Element].AddError()
		}
	}
}

func (r *AccessibilityReport) Save() {
	for Element, Summary := range r.Summary {
		log.Printf("%d %s tested with %d errors and %d asserts", Summary.Total, Element, Summary.Errors, Summary.Pass)
	}
}

func NewAccessibilityReport(url string, title string, lang string) AccessibilityReport {
	return AccessibilityReport{URL: url, Title: title, Lang: lang}
}
