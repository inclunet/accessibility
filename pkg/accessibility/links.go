package accessibility

type Links struct {
	Element
}

func (l *Links) Check() ([]AccessibilityCheck, bool) {
	accessibleText, ok := l.GetAccessibleText()

	if !ok {
		accessibilityChecks := Check(l.Selection.Children(), []AccessibilityCheck{})
		for _, accessibilityCheck := range accessibilityChecks {
			return l.AddViolation(accessibilityCheck.Violation, false)
		}

		return l.AddViolation("emag-3.5.3", false)
	}

	return NewAccessibleTextCheck(l.AccessibilityChecks, l.AccessibilityCheck).SetMaxLength(200, "too-long-text").Check(accessibleText)
}
