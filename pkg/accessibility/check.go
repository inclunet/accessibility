package accessibility

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

func DeepCheck(s *goquery.Selection, accessibilityReport report.AccessibilityReport) (int, bool, string, error) {
	A, Pass, Description, err := NewElementCheck(s, accessibilityReport)
	if err != nil {
		s.Each(func(i int, s *goquery.Selection) {
			A, Pass, Description, err = DeepCheck(s.Children(), accessibilityReport)
		})
	}
	return A, Pass, Description, err
}

func GetElementInterface(elementName string) (func(*goquery.Selection, report.AccessibilityReport) Accessibility, error) {
	checkList := map[string]func(*goquery.Selection, report.AccessibilityReport) Accessibility{
		"a":      NewLinkCheck,
		"button": NewButtonCheck,
		"h1":     NewHeaderCheck,
		"h2":     NewHeaderCheck,
		"h3":     NewHeaderCheck,
		"h4":     NewHeaderCheck,
		"h5":     NewHeaderCheck,
		"h6":     NewHeaderCheck,
		"input":  NewInputCheck,
		"img":    NewImageCheck,
	}
	if elementInterface, ok := checkList[elementName]; ok {
		return elementInterface, nil
	}
	return nil, errors.New("no evaluator available to this element type")
}

func NewElementCheck(s *goquery.Selection, accessibilityReport report.AccessibilityReport) (int, bool, string, error) {
	elementInterface, err := GetElementInterface(goquery.NodeName(s))

	if err != nil {
		return 1, false, "", err
	}

	accessibilityInterface := elementInterface(s, accessibilityReport)
	a, pass, description := accessibilityInterface.Check()
	return a, pass, description, nil
}
