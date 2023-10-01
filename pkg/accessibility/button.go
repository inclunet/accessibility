package accessibility

type Buttons struct {
	Element
}

func (b *Buttons) Check() ([]AccessibilityCheck, bool) {
	accessibleText, ok := b.GetAccessibleText()

	if !ok {
		accessibilityChecks := Check(b.Selection.Children(), []AccessibilityCheck{})
		for _, accessibilityCheck := range accessibilityChecks {
			return b.AddViolation(accessibilityCheck.Violation, false)
		}

		return b.AddViolation("emag-6.1.1", false)
	}

	return NewAccessibleTextCheck(b.AccessibilityChecks, b.AccessibilityCheck).SetMaxLength(120, "too-long-text").Check(accessibleText)
}
