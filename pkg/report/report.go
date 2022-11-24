package report

import (
	"html/template"
	"os"
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
	Summary map[string]AccessibilitySummary
	Checks  []AccessibilityCheck
}

func (r *AccessibilityReport) AddCheck(Element string, a int, pass bool, description string, Html string) {
	r.Checks = append(r.Checks, AccessibilityCheck{Element: Element, A: a, Pass: pass, Description: description, Html: Html})
	r.UpdateSummary(Element, pass)
}

func (r *AccessibilityReport) UpdateSummary(Element string, Pass bool) {
	Summary, _ := r.Summary[Element]
	Summary.Update(Pass)
	r.Total = r.Total + 1
	if Pass {
		r.Pass = r.Pass + 1
	} else {
		r.Errors = r.Errors + 1
	}
	r.Summary[Element] = Summary
}

func (r *AccessibilityReport) GenerateSummary() {
	for _, check := range r.Checks {
		r.UpdateSummary(check.Element, check.Pass)
	}
}

func (r *AccessibilityReport) Save(report string) error {
	Template := template.Must(template.New("model.html").ParseFiles("model.html"))

	File, err := os.Create(report)

	if err != nil {
		return err
	}

	defer File.Close()

	err = Template.Execute(File, r)

	if err != nil {
		return err
	}

	return nil
}

func NewAccessibilityReport(url string, title string, lang string) AccessibilityReport {
	Report := AccessibilityReport{
		URL:     url,
		Title:   title,
		Lang:    lang,
		Summary: make(map[string]AccessibilitySummary),
	}
	return Report
}
