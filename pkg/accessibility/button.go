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

	return NewAccessibleTextCheck(accessibilityCheck).SetMaxLength(120, "too-long-text").Check(accessibleText)
}
