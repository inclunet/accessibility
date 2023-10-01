package accessibility

type AccessibleText struct {
	accessibilityCheck  AccessibilityCheck
	accessibilityChecks []AccessibilityCheck
	maxLength           int
	minLength           int
	violationEmpty      string
	violationMaxLength  string
	violationMinLength  string
}

func (a *AccessibleText) AddViolation(violation string, pass bool) ([]AccessibilityCheck, bool) {
	a.accessibilityChecks = append(a.accessibilityChecks, a.accessibilityCheck.SetViolation(violation))
	return a.accessibilityChecks, pass
}

func (a *AccessibleText) Check(accessibleText string) ([]AccessibilityCheck, bool) {
	if accessibleText == "" {
		return a.AddViolation(a.violationEmpty, false)
	}

	if len(accessibleText) < a.minLength {
		return a.AddViolation(a.violationMinLength, false)
	}

	if len(accessibleText) > a.maxLength {
		return a.AddViolation(a.violationMaxLength, false)
	}

	return a.AddViolation("pass", true)
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

func NewAccessibleTextCheck(accessibilityChecks []AccessibilityCheck, accessibilityCheck AccessibilityCheck) *AccessibleText {
	return &AccessibleText{
		accessibilityCheck:  accessibilityCheck,
		accessibilityChecks: accessibilityChecks,
		maxLength:           120,
		minLength:           3,
		violationEmpty:      "emag-1.2.3",
		violationMaxLength:  "too-long-text",
		violationMinLength:  "too-short-text",
	}
}
