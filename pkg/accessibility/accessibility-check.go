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

func (c *AccessibilityCheck) SetViolation(violation string) AccessibilityCheck {
	c.Violation = violation
	return *c
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
