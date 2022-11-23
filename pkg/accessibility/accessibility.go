package accessibility

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type HtmlElement interface {
	AriaHidden() bool
	AriaLabel() (string, bool)
	Check(s *goquery.Selection, Report *report.AccessibilityReport) (int, bool, string)
	Role() (string, bool)
}

type Element struct {
	Selection *goquery.Selection
}

func (e *Element) AriaHidden() bool {
	if value, ok := e.Selection.Attr("aria-hidden"); ok {
		if value == "true" {
			return true
		}
	}
	return false
}

func (e *Element) AriaLabel() (string, bool) {
	if value, ok := e.Selection.Attr("aria-label"); ok {
		if value != "" {
			return value, true
		}
	}
	return "", false
}

func (e *Element) Role() bool {
	if value, ok := e.Selection.Attr("role"); ok {
		if value == "true" {
			return true
		}
	}
	return false
}
