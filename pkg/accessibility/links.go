package accessibility

type Links struct {
	Element
}

func (l *Links) Check() AccessibilityCheck {
	accessibilityCheck := l.NewAccessibilityCheck("pass")

	if l.AriaHidden() {
		return accessibilityCheck.SetViolation("aria-hidden")
	}

	accessibleText, ok := l.AccessibleText()

	if !ok {
		if accessibilityCheck, err := l.DeepCheck(l.Selection.Children(), l.AccessibilityChecks); err == nil {
			return accessibilityCheck
		}

		return accessibilityCheck.SetViolation("emag-3.5.3")
	}

	if l.CheckTooShortText(accessibleText) {
		return accessibilityCheck.SetViolation("too-short-text")
	}

	if l.CheckTooLongText(accessibleText, 200) {
		return accessibilityCheck.SetViolation("too-long-text")
	}

	return accessibilityCheck
}
