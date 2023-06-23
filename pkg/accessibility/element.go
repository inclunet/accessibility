package accessibility

import (
	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	AccessibilityChecks []AccessibilityCheck
	Selection           *goquery.Selection
	useAlternativeText  bool
	useAriaHidden       bool
	useAriaLabel        bool
	useElementText      bool
	useTitle            bool
}

func (e *Element) Check() AccessibilityCheck {
	return AccessibilityCheck{}
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

func (e *Element) AriaHidden() bool {
	if value, ok := e.Selection.Attr("aria-hidden"); ok && value == "true" {
		return true
	}
	return false
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

func (e *Element) CheckAccessibleText(accessibilityCheck AccessibilityCheck) AccessibilityCheck {

	return accessibilityCheck
}

func (e *Element) NewAccessibilityCheck(violation string) AccessibilityCheck {
	htmlElement, _ := goquery.OuterHtml(e.Selection)
	return NewAccessibilityCheck(goquery.NodeName(e.Selection), htmlElement, violation)
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

func NewElementCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) (AccessibilityCheck, error) {
	accessibilityInterface, err := GetElementInterface(goquery.NodeName(s))

	if err != nil {
		return AccessibilityCheck{}, err
	}

	accessibilityInterface.SetSelection(s).SetAccessibilityChecks(accessibilityChecks).SetUseAlternativeText(true).SetUseAriaLabel(true).SetUseElementText(true).SetUseTitle(true).SetUseAriaHidden(true)

	return accessibilityInterface.Check(), nil
}
