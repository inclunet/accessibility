package accessibility

type AmpImg struct {
	Images
}

func (a *AmpImg) Check() AccessibilityCheck {
	AccessibilityCheck := a.Images.Check()

	if !AccessibilityCheck.Pass {
		deepAccessibilityCheck, err := a.DeepCheck(a.Selection.Children(), a.AccessibilityChecks)

		if err == nil {
			return deepAccessibilityCheck
		}

		AccessibilityCheck.Description = err.Error()
	}

	return AccessibilityCheck
}
