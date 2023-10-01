package accessibility

import "html/template"

type AccessibilityCheck struct {
	Element     string
	A           int
	Error       bool
	Warning     bool
	Description string
	Html        template.HTML
	Line        int
	Solution    string
	Text        string
	Violation   string
}

func (a *AccessibilityCheck) SetElement(element string) *AccessibilityCheck {
	a.Element = element
	return a
}

func (a *AccessibilityCheck) SetHtml(html string) *AccessibilityCheck {
	a.Html = template.HTML(html)
	a.SetText(html)
	return a
}

func (a *AccessibilityCheck) SetText(text string) *AccessibilityCheck {
	a.Text = text
	return a
}

func (a *AccessibilityCheck) SetViolation(violation string) AccessibilityCheck {
	a.Violation = violation
	return *a
}

func NewAccessibilityCheck(element string, htmlElement string, violation string) AccessibilityCheck {
	return AccessibilityCheck{
		Description: "",
		Element:     element,
		Error:       false,
		Html:        template.HTML(htmlElement),
		Solution:    "",
		Text:        htmlElement,
		Violation:   violation,
		Warning:     false,
	}
}
