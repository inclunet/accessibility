package accessibility

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

type Accessibility interface {
	GetAlternativeText() (string, bool)
	GetAriaLabel() (string, bool)
	GetTitle() (string, bool)
	AriaHidden() bool
	Check() AccessibilityCheck
	DeepCheck(*goquery.Selection, []AccessibilityCheck) (AccessibilityCheck, error)
	NewAccessibilityCheck(string) AccessibilityCheck
	Role() (string, bool)
	SetAccessibilityChecks(accessibilityChecks []AccessibilityCheck) Accessibility
	SetSelection(s *goquery.Selection) Accessibility
	SetUseAlternativeText(bool) Accessibility
	SetUseAriaHidden(bool) Accessibility
	SetUseAriaLabel(bool) Accessibility
	SetUseElementText(bool) Accessibility
	SetUseTitle(bool) Accessibility
}

func AfterCheck(accessibilityChecks []AccessibilityCheck) []AccessibilityCheck {
	newChecks := []AccessibilityCheck{}

	if HeaderUnavailable(accessibilityChecks) {
		newChecks = append(newChecks, NewAccessibilityCheck("h1", "", "emag-1.3.1"))
	}

	if HeaderInvalidOrdenation(accessibilityChecks) {
		newChecks = append(newChecks, NewAccessibilityCheck("h2", "", "emag-1.3.3"))
	}

	if HeaderMainUnavailable(accessibilityChecks) {
		newChecks = append(newChecks, NewAccessibilityCheck("h1", "", "emag-1.3.4"))
	}

	return newChecks
}

func GetElementInterface(elementName string) (Accessibility, error) {
	checkList := map[string]Accessibility{
		"a": &Links{},
		//"amp-img": &AmpImg{},
		"button": &Buttons{},
		"h1":     &Headers{},
		"h2":     &Headers{},
		"h3":     &Headers{},
		"h4":     &Headers{},
		"h5":     &Headers{},
		"h6":     &Headers{},
		"input":  &Inputs{},
		"img":    &Images{},
	}

	if elementInterface, ok := checkList[elementName]; ok {
		return elementInterface, nil
	}

	return nil, errors.New("no evaluator available to this element type")
}
