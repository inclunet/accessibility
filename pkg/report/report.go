package report

import (
	"log"
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

func (r *AccessibilityReport) AddCheck(Element string, a int, pass bool, description string, Html string) {
	r.Checks = append(r.Checks, AccessibilityCheck{Element: Element, A: a, Pass: pass, Description: description, Html: Html})
	r.UpdateSummary(Element, pass)
}

func (r *AccessibilityReport) UpdateSummary(Element string, Pass bool) {
	if _, ok := r.Summary[Element]; !ok {
		r.Summary[Element] = NewSummary()
	}
	if Pass {
		r.Summary[Element].AddPass()
	} else {
		r.Summary[Element].AddError()
	}
}

func (r *AccessibilityReport) GenerateSummary() {
	for _, check := range r.Checks {
		r.UpdateSummary(check.Element, check.Pass)
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
