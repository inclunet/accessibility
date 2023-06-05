package accessibility

import (
	"html/template"

	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	Selection           *goquery.Selection
	AccessibilityChecks []AccessibilityCheck
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

func (e *Element) DeepCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) (AccessibilityCheck, error) {
	accessibilityCheck, err := NewElementCheck(s, accessibilityChecks)

	if err != nil {
		s.Each(func(i int, s *goquery.Selection) {
			accessibilityCheck, err = e.DeepCheck(s.Children(), accessibilityChecks)
		})
	}
	return accessibilityCheck, err
}

func (e *Element) CheckTooLongText(accessibilityText string, maxLength int) bool {
	return len(accessibilityText) > maxLength
}

func (e *Element) CheckTooShortText(accessibilityText string) bool {
	return len(accessibilityText) < 3
}

func (e *Element) NewAccessibilityCheck(a int, description string) AccessibilityCheck {
	htmlElement, _ := goquery.OuterHtml(e.Selection)
	return AccessibilityCheck{
		Element:     goquery.NodeName(e.Selection),
		A:           a,
		Pass:        false,
		Warning:     false,
		Description: description,
		Html:        template.HTML(htmlElement),
		Text:        htmlElement,
	}
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

func NewElementCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) (AccessibilityCheck, error) {
	elementInterface, err := GetElementInterface(goquery.NodeName(s))

	if err != nil {
		return AccessibilityCheck{}, err
	}

	accessibilityInterface := elementInterface(s, accessibilityChecks)
	return accessibilityInterface.Check(), nil
}
