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

	return NewAccessibleTextCheck(accessibilityCheck).SetMaxLength(200, "too-long-text").Check(accessibleText)
}
