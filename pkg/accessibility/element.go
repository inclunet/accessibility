package accessibility

import (
	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	AccessibilityCheck  AccessibilityCheck
	AccessibilityChecks []AccessibilityCheck
	Selection           *goquery.Selection
	useAlternativeText  bool
	useAriaHidden       bool
	useAriaLabel        bool
	useElementText      bool
	useTitle            bool
}

func (e *Element) AddViolation(violation string, pass bool) ([]AccessibilityCheck, bool) {
	e.AccessibilityChecks = append(e.AccessibilityChecks, e.AccessibilityCheck.SetViolation(violation))
	return e.AccessibilityChecks, pass
}

func (e *Element) AfterCheck() ([]AccessibilityCheck, bool) {
	return e.AccessibilityChecks, false
}

func (e *Element) BeforeCheck() ([]AccessibilityCheck, bool) {
	return e.AccessibilityChecks, false
}

func (e *Element) Check() ([]AccessibilityCheck, bool) {
	return e.AccessibilityChecks, false
}

func (e *Element) GetAccessibleText() (string, bool) {
	if accessibleText, ok := e.GetAriaLabel(); ok {
		return accessibleText, ok
	}

	if accessibleText, ok := e.GetAlternativeText(); ok {
		return accessibleText, ok
	}

	if accessibleText := e.Selection.Text(); accessibleText != "" && e.useElementText {
		return accessibleText, true
	}

	if accessibleText, ok := e.GetTitle(); ok {
		return accessibleText, ok
	}

	return "", false
}

func (e *Element) GetAlternativeText() (string, bool) {
	if accessibleText, ok := e.Selection.Attr("alt"); ok && accessibleText != "" && e.useAlternativeText {
		return accessibleText, true
	}
	return "", false
}

func (e *Element) GetAriaLabel() (string, bool) {
	if value, ok := e.Selection.Attr("aria-label"); ok && value != "" && e.useAriaLabel {
		return value, true
	}
	return "", false
}

func (e *Element) GetTitle() (string, bool) {
	if accessibleText, ok := e.Selection.Attr("title"); ok && accessibleText != "" && e.useTitle {
		return accessibleText, true
	}
	return "", false
}

func (e *Element) IsAriaHidden() ([]AccessibilityCheck, bool) {
	if value, ok := e.Selection.Attr("aria-hidden"); ok && value == "true" && e.useAriaHidden {
		return e.AddViolation("aria-hidden", true)
	}
	return e.AddViolation("pass", false)
}

func (e *Element) CheckAccessibleText(accessibilityCheck AccessibilityCheck) AccessibilityCheck {

	return accessibilityCheck
}

func (e *Element) NewAccessibilityCheck() Accessibility {
	html, _ := goquery.OuterHtml(e.Selection)
	e.AccessibilityCheck.SetElement(goquery.NodeName(e.Selection)).SetHtml(html)
	return e
}

func (e *Element) Role() (string, bool) {
	if role, ok := e.Selection.Attr("role"); ok && role != "" {
		return role, ok
	}
	return "", false
}

func (e *Element) SetAccessibilityChecks(accessibilityChecks []AccessibilityCheck) Accessibility {
	e.AccessibilityChecks = accessibilityChecks
	return e
}

func (e *Element) SetSelection(s *goquery.Selection) Accessibility {
	e.Selection = s
	return e
}

func (e *Element) SetUseAlternativeText(useAlternativeText bool) Accessibility {
	e.useAlternativeText = useAlternativeText
	return e
}

func (e *Element) SetUseAriaHidden(useAriaHidden bool) Accessibility {
	e.useAriaHidden = useAriaHidden
	return e
}

func (e *Element) SetUseAriaLabel(useAriaLabel bool) Accessibility {
	e.useAriaLabel = useAriaLabel
	return e
}

func (e *Element) SetUseElementText(useElementText bool) Accessibility {
	e.useElementText = useElementText
	return e
}

func (e *Element) SetUseTitle(useTitle bool) Accessibility {
	e.useTitle = useTitle
	return e
}

func NewElementCheck(s *goquery.Selection) (Accessibility, error) {
	accessibilityInterface, err := GetElementInterface(goquery.NodeName(s))

	if err != nil {
		return nil, err
	}

	return accessibilityInterface.SetSelection(s).NewAccessibilityCheck(), nil
}
