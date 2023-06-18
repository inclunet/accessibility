package accessibility

type Images struct {
	Element
}

func (i *Images) Check() AccessibilityCheck {
	accessibilityCheck := i.NewAccessibilityCheck("pass")

	if i.AriaHidden() {
		return i.FindViolation(accessibilityCheck, "aria-hidden")
	}

	accessibleText, ok := i.AccessibleText()

	if !ok {
		return i.FindViolation(accessibilityCheck, "emag-3.6.2")
	}

	if i.CheckTooShortText(accessibleText) {
		return i.FindViolation(accessibilityCheck, "too-short-text")
	}

	if i.CheckTooLongText(accessibleText, 240) {
		return i.FindViolation(accessibilityCheck, "too-long-text")
	}

	return accessibilityCheck
}
