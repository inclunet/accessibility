package accessibility

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

func NewCheck(s *goquery.Selection, accessibilityReport *report.AccessibilityReport) (int, bool, string, error) {
	accessibilityCheckList := map[string]func(*goquery.Selection, *report.AccessibilityReport) Accessibility{
		"a":   NewLinkCheck,
		"h1":  NewHeaderCheck,
		"h2":  NewHeaderCheck,
		"h3":  NewHeaderCheck,
		"h4":  NewHeaderCheck,
		"h5":  NewHeaderCheck,
		"h6":  NewHeaderCheck,
		"img": NewImageCheck,
	}
	elementName := goquery.NodeName(s)
	if elementInterface, ok := accessibilityCheckList[elementName]; ok {
		accessibilityInterface := elementInterface(s, accessibilityReport)
		a, pass, description := accessibilityInterface.Check()
		return a, pass, description, nil
	}
	return 1, false, "", errors.New("no defined evaluator function for this html element")
}
