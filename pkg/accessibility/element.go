package accessibility

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type Element struct {
	Selection           *goquery.Selection
	AccessibilityReport report.AccessibilityReport
}

func (e *Element) AlternativeText() (string, bool) {
	if accessibleText, ok := e.Selection.Attr("alt"); ok && accessibleText != "" {
		return accessibleText, true
	}
	return "", false
}

func (e *Element) AriaHidden() bool {
	if value, ok := e.Selection.Attr("aria-hidden"); ok && value == "true" {
		return true
	}
	return false
}

func (e *Element) AriaLabel() (string, bool) {
	if value, ok := e.Selection.Attr("aria-label"); ok && value != "" {
		return value, true
	}
	return "", false
}

func (e *Element) AccessibleText() (string, bool) {
	if accessibleText, ok := e.AriaLabel(); ok {
		return accessibleText, ok
	}

	if accessibleText, ok := e.AlternativeText(); ok {
		return accessibleText, ok
	}

	if accessibleText := e.Selection.Text(); accessibleText != "" {
		return accessibleText, true
	}

	if accessibleText, ok := e.Title(); ok {
		return accessibleText, ok
	}

	return "", false
}

func (e *Element) Role() (string, bool) {
	if role, ok := e.Selection.Attr("role"); ok && role != "" {
		return role, ok
	}
	return "", false
}

func (e *Element) Title() (string, bool) {
	if accessibleText, ok := e.Selection.Attr("title"); ok && accessibleText != "" {
		return accessibleText, true
	}
	return "", false
}
