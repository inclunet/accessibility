package accessibility

type Images struct {
	Element
}

func (i *Images) Check() AccessibilityCheck {
	accessibilityCheck := i.NewAccessibilityCheck("pass")

	if i.AriaHidden() {
		return accessibilityCheck.SetViolation("aria-hidden")
	}

	accessibleText, ok := i.AccessibleText()

	if !ok {
		return accessibilityCheck.SetViolation("emag-3.6.2")
	}

	return NewAccessibleTextCheck(accessibilityCheck).SetMaxLength(240, "too-long-text").Check(accessibleText)
}
