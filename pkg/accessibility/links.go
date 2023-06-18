package accessibility

type Links struct {
	Element
}

func (l *Links) Check() AccessibilityCheck {
	accessibilityCheck := l.NewAccessibilityCheck("pass")

	if l.AriaHidden() {
		return l.FindViolation(accessibilityCheck, "aria-hidden")
	}

	accessibleText, ok := l.AccessibleText()

	if !ok {
		if accessibilityCheck, err := l.DeepCheck(l.Selection.Children(), l.AccessibilityChecks, l.AccessibilityRules); err == nil {
			return accessibilityCheck
		}

		return l.FindViolation(accessibilityCheck, "emag-3.5.3")
	}

	if l.CheckTooShortText(accessibleText) {
		return l.FindViolation(accessibilityCheck, "too-short-text")
	}

	if l.CheckTooLongText(accessibleText, 200) {
		return l.FindViolation(accessibilityCheck, "too-long-text")
	}

	return accessibilityCheck
}
