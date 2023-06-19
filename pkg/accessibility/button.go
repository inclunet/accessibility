package accessibility

type Buttons struct {
	Element
}

func (b *Buttons) Check() AccessibilityCheck {
	accessibilityCheck := b.NewAccessibilityCheck("pass")

	if b.AriaHidden() {
		return b.FindViolation(accessibilityCheck, "aria-hidden")
	}

	accessibleText, ok := b.AccessibleText()

	if !ok {
		if accessibilityCheck, err := b.DeepCheck(b.Selection.Children(), b.AccessibilityChecks, b.AccessibilityRules); err == nil {
			return accessibilityCheck
		}

		return b.FindViolation(accessibilityCheck, "emag-6.1.1")
	}

	if b.CheckTooShortText(accessibleText) {
		return b.FindViolation(accessibilityCheck, "too-short-text")
	}

	if b.CheckTooLongText(accessibleText, 120) {
		return b.FindViolation(accessibilityCheck, "too-long-text")
	}

	return accessibilityCheck
}
