package accessibility

type AmpImg struct {
	Images
}

func (a *AmpImg) Check() AccessibilityCheck {
	AccessibilityCheck := a.Images.Check()

	if AccessibilityCheck.Error {
		if accessibilityCheck, err := a.DeepCheck(a.Selection.Children(), a.AccessibilityChecks, a.AccessibilityRules); err == nil {
			return accessibilityCheck
		}

		return a.FindViolation(AccessibilityCheck, "emag-3.6.2")
	}

	return AccessibilityCheck
}
