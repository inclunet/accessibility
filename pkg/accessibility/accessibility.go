package accessibility

import (
	"errors"
	"html/template"

	"github.com/PuerkitoBio/goquery"
)

type AccessibilityRule struct {
	A           int
	Description string
	Error       bool
	Solution    string
	Warning     bool
}

type AccessibilityCheck struct {
	Element     string
	A           int
	Pass        bool
	Warning     bool
	Description string
	Html        template.HTML
	Line        int
	Solution    string
	Text        string
}

type Accessibility interface {
	AlternativeText() (string, bool)
	AriaHidden() bool
	AriaLabel() (string, bool)
	Check() AccessibilityCheck
	DeepCheck(*goquery.Selection, []AccessibilityCheck) (AccessibilityCheck, error)
	NewAccessibilityCheck(int, string) AccessibilityCheck
	Role() (string, bool)
	SetAccessibilityChecks(accessibilityChecks []AccessibilityCheck)
	SetAccessibilityRules(accessibilityRules *map[string]AccessibilityRule)
	SetSelection(s *goquery.Selection)
}

func GetElementInterface(elementName string) (Accessibility, error) {
	checkList := map[string]Accessibility{
		"a":       &Links{},
		"amp-img": &AmpImg{},
		"button":  &Buttons{},
		"h1":      &Headers{},
		"h2":      &Headers{},
		"h3":      &Headers{},
		"h4":      &Headers{},
		"h5":      &Headers{},
		"h6":      &Headers{},
		"input":   &Inputs{},
		"img":     &Images{},
	}

	if elementInterface, ok := checkList[elementName]; ok {
		return elementInterface, nil
	}

	return nil, errors.New("no evaluator available to this element type")
}
