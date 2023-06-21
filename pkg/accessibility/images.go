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

	if i.CheckTooShortText(accessibleText) {
		return accessibilityCheck.SetViolation("too-short-text")
	}

	if i.CheckTooLongText(accessibleText, 240) {
		return accessibilityCheck.SetViolation("too-long-text")
	}

	return accessibilityCheck
}
