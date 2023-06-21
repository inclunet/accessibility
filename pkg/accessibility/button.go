package accessibility

type Buttons struct {
	Element
}

func (b *Buttons) Check() AccessibilityCheck {
	accessibilityCheck := b.NewAccessibilityCheck("pass")

	if b.AriaHidden() {
		return accessibilityCheck.SetViolation("aria-hidden")
	}

	accessibleText, ok := b.AccessibleText()

	if !ok {
		if accessibilityCheck, err := b.DeepCheck(b.Selection.Children(), b.AccessibilityChecks); err == nil {
			return accessibilityCheck
		}

		return accessibilityCheck.SetViolation("emag-6.1.1")
	}

	if b.CheckTooShortText(accessibleText) {
		return accessibilityCheck.SetViolation("too-short-text")
	}

	if b.CheckTooLongText(accessibleText, 120) {
		return accessibilityCheck.SetViolation("too-long-text")
	}

	return accessibilityCheck
}
