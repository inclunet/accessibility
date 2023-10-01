package accessibility

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/PuerkitoBio/goquery"
)

type Accessibility interface {
	AddViolation(string, bool) ([]AccessibilityCheck, bool)
	AfterCheck() ([]AccessibilityCheck, bool)
	BeforeCheck() ([]AccessibilityCheck, bool)
	Check() ([]AccessibilityCheck, bool)
	GetAlternativeText() (string, bool)
	GetAriaLabel() (string, bool)
	GetTitle() (string, bool)
	IsAriaHidden() ([]AccessibilityCheck, bool)
	NewAccessibilityCheck() Accessibility
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

func Check(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) []AccessibilityCheck {
	s.Each(func(i int, s *goquery.Selection) {
		pass := false
		newElement, err := NewElementCheck(s)

		if err == nil {
			newElement.SetAccessibilityChecks(accessibilityChecks).SetUseAlternativeText(true).SetUseAriaHidden(true).SetUseAriaLabel(true).SetUseElementText(true).SetUseTitle(true)
			accessibilityChecks, pass = newElement.BeforeCheck()

			if !pass {
				accessibilityChecks, pass = newElement.IsAriaHidden()
			}

			if !pass {
				fmt.Println(goquery.NodeName(s) + ": " + reflect.TypeOf(newElement).Name())

				accessibilityChecks, pass = newElement.Check()
			}

			if !pass {
				accessibilityChecks, _ = newElement.AfterCheck()
			}
		}

		accessibilityChecks = Check(s.Children(), accessibilityChecks)
	})

	return accessibilityChecks
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
