package accessibility

type AmpImg struct {
	Images
}

func (a *AmpImg) Check() AccessibilityCheck {
	accessibilityCheck := a.Images.Check()

	if accessibilityCheck.Error {
		if accessibilityCheck, err := a.DeepCheck(a.Selection.Children(), a.AccessibilityChecks); err == nil {
			return accessibilityCheck
		}

		return accessibilityCheck.SetViolation("emag-3.6.2")
	}

	return accessibilityCheck
}
