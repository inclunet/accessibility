package accessibility

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

type AccessibilityCheck struct {
	Element     string
	A           int
	Pass        bool
	Warning     bool
	Description string
	Line        int
	Html        string
}

type Accessibility interface {
	AlternativeText() (string, bool)
	AriaHidden() bool
	AriaLabel() (string, bool)
	Check() AccessibilityCheck
	DeepCheck(*goquery.Selection, []AccessibilityCheck) (AccessibilityCheck, error)
	NewAccessibilityCheck(int, string) AccessibilityCheck
	Role() (string, bool)
}

func GetElementInterface(elementName string) (func(*goquery.Selection, []AccessibilityCheck) Accessibility, error) {
	checkList := map[string]func(*goquery.Selection, []AccessibilityCheck) Accessibility{
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
