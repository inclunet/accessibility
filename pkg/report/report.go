package report

import (
	"html"
	"strings"

	"github.com/inclunet/accessibility/pkg/accessibility"
)

type AccessibilityReport struct {
	Checks         []accessibility.AccessibilityCheck
	Domain         string
	Html           string
	HtmlReportPath string
	JsonReportPath string
	ReportFile     string
	Summary        map[string]AccessibilitySummary
	Title          string
	Url            string
}

func (r *AccessibilityReport) AddCheck(accessibilityCheck accessibility.AccessibilityCheck) {
	r.Checks = append(r.Checks, r.GetLineNumber(accessibilityCheck))
}

func (r *AccessibilityReport) FindViolation(accessibilityViolations map[string]accessibility.AccessibilityViolation, accessibilityCheck accessibility.AccessibilityCheck) accessibility.AccessibilityCheck {
	if accessibilityViolation, ok := accessibilityViolations[accessibilityCheck.Violation]; ok {
		accessibilityCheck.A = accessibilityViolation.A
		accessibilityCheck.Description = accessibilityViolation.Description
		accessibilityCheck.Solution = accessibilityViolation.Solution
		accessibilityCheck.Error = accessibilityViolation.Error
		accessibilityCheck.Warning = accessibilityViolation.Warning
	}

	return accessibilityCheck
}

func (r *AccessibilityReport) FindViolations(accessibilityViolations map[string]accessibility.AccessibilityViolation) {
	for i, accessibilityCheck := range r.Checks {
		r.Checks[i] = r.FindViolation(accessibilityViolations, accessibilityCheck)
	}
}

func (r *AccessibilityReport) UpdateSummary(accessibilityCheck accessibility.AccessibilityCheck) {
	total := r.Summary["total"]
	total.Update(accessibilityCheck)
	r.Summary["total"] = total
	Summary := r.Summary[accessibilityCheck.Element]
	Summary.Update(accessibilityCheck)
	r.Summary[accessibilityCheck.Element] = Summary
}

func (r *AccessibilityReport) GenerateSummary() {
	r.Summary = make(map[string]AccessibilitySummary)
	for _, accessibilityCheck := range r.Checks {
		r.UpdateSummary(accessibilityCheck)
	}
}

func (r *AccessibilityReport) GetLineNumber(accessibilityCheck accessibility.AccessibilityCheck) accessibility.AccessibilityCheck {
	text := html.UnescapeString(accessibilityCheck.Text)
	offset := strings.Index(r.Html, text)

	if offset > 0 {
		accessibilityCheck.Line = strings.Count(r.Html[:offset], "\n")
	}

	return accessibilityCheck
}

func NewAccessibilityReport(url string, reportFile string, lang string, reportPath string, title string) AccessibilityReport {
	return AccessibilityReport{
		Url:        url,
		Title:      title,
		ReportFile: reportFile,
	}
}
