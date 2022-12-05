package accessibility

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type Headers struct {
	Element
}

func (h *Headers) isHeader(elementName string) bool {
	switch elementName {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return true
	default:
		return false
	}
}

func (h *Headers) isIncorrectLevel(last int, actual int) bool {
	return actual > last+1
}

func (h *Headers) isTooLongHeader() bool {
	if accessibleText, ok := h.AccessibleText(); ok && len(accessibleText) > 80 {
		return true
	}
	return false
}

func (h *Headers) CheckHierarchy(elementName string) bool {
	headerLevel := 0
	for _, check := range h.AccessibilityReport.Checks {
		if h.isHeader(check.Element) {
			actualLevel, _ := strconv.Atoi(strings.TrimPrefix(check.Element, "h"))
			if h.isIncorrectLevel(headerLevel, actualLevel) {
				return false
			}
			headerLevel = actualLevel
		}
	}
	actualLevel, _ := strconv.Atoi(strings.TrimPrefix(elementName, "h"))
	if h.isIncorrectLevel(headerLevel, actualLevel) {
		return false
	} else {
		return true
	}
}

func (h *Headers) Check() (int, bool, string) {
	switch {
	case h.CheckHierarchy(goquery.NodeName(h.Selection)):
		return 1, true, "Please check if this header is following the headers hierarchy"
	case h.isTooLongHeader():
		return 1, false, "Please check if this header is too longh, too long headers is not presented at a single line and is possibly a bad practice."
	default:
		return 1, true, "The headers are ok"
	}
}

func NewHeaderCheck(s *goquery.Selection, accessibilityReport *report.AccessibilityReport) Accessibility {
	accessibilityInterface := new(Headers)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityReport = accessibilityReport
	return accessibilityInterface
}
