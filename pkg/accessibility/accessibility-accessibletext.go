package accessibility

type AccessibleText struct {
	accessibilityCheck AccessibilityCheck
	maxLength          int
	minLength          int
	violationEmpty     string
	violationMaxLength string
	violationMinLength string
}

func (a *AccessibleText) Check(accessibleText string) AccessibilityCheck {
	if accessibleText == "" {
		return a.accessibilityCheck.SetViolation(a.violationEmpty)
	}

	if len(accessibleText) < a.minLength {
		return a.accessibilityCheck.SetViolation(a.violationMinLength)
	}

	if len(accessibleText) > a.maxLength {
		return a.accessibilityCheck.SetViolation(a.violationMaxLength)
	}

	return a.accessibilityCheck
}

func (a *AccessibleText) SetEmptyViolation(accessibilityViolation string) *AccessibleText {
	a.violationEmpty = accessibilityViolation
	return a
}

func (a *AccessibleText) SetMaxLength(maxLength int, accessibilityViolation string) *AccessibleText {
	a.maxLength = maxLength
	a.violationMaxLength = accessibilityViolation
	return a
}

func (a *AccessibleText) SetMinLength(MinLength int, accessibilityViolation string) *AccessibleText {
	a.minLength = MinLength
	a.violationMinLength = accessibilityViolation
	return a
}

func NewAccessibleTextCheck(accessibilityCheck AccessibilityCheck) *AccessibleText {
	return &AccessibleText{
		accessibilityCheck: accessibilityCheck,
		maxLength:          120,
		minLength:          3,
		violationEmpty:     "emag-1.2.3",
		violationMaxLength: "too-long-text",
		violationMinLength: "too-short-text",
	}
}
