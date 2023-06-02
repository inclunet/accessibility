package report

import (
	"github.com/inclunet/accessibility/pkg/accessibility"
)

type AccessibilityReport struct {
	Checks     []accessibility.AccessibilityCheck
	Domain     string
	Errors     int
	Html       string
	Pass       int
	ReportFile string
	Summary    map[string]AccessibilitySummary
	Title      string
	Total      int
	Url        string
}

func (r *AccessibilityReport) AddCheck(accessibilityCheck accessibility.AccessibilityCheck) {
	r.Checks = append(r.Checks, accessibilityCheck)
	r.UpdateSummary(accessibilityCheck)
}

func (r *AccessibilityReport) UpdateSummary(accessibilityCheck accessibility.AccessibilityCheck) {
	Summary, _ := r.Summary[accessibilityCheck.Element]
	Summary.Update(accessibilityCheck.Pass)
	r.Total = r.Total + 1
	if accessibilityCheck.Pass {
		r.Pass = r.Pass + 1
	} else {
		r.Errors = r.Errors + 1
	}
	r.Summary[accessibilityCheck.Element] = Summary
}

func (r *AccessibilityReport) GenerateSummary() {
	for _, accessibilityCheck := range r.Checks {
		r.UpdateSummary(accessibilityCheck)
	}
}

func (r *AccessibilityReport) NewSummary() {
	r.Summary = make(map[string]AccessibilitySummary)
}

func NewAccessibilityReport(url string, reportFile string, lang string, reportPath string, title string) AccessibilityReport {
	return AccessibilityReport{
		Url:        url,
		Title:      title,
		ReportFile: reportFile,
		Summary:    make(map[string]AccessibilitySummary),
	}
}
