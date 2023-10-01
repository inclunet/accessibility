package accessibility

type Images struct {
	Element
}

func (i *Images) Check() ([]AccessibilityCheck, bool) {
	accessibleText, ok := i.GetAccessibleText()

	if !ok {
		return i.AddViolation("emag-3.6.2", false)
	}

	return NewAccessibleTextCheck(i.AccessibilityChecks, i.AccessibilityCheck).SetMaxLength(240, "too-long-text").Check(accessibleText)
}
