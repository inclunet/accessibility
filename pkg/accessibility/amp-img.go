package accessibility

import "github.com/PuerkitoBio/goquery"

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

func NewAmpImageCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) Accessibility {
	accessibilityInterface := new(AmpImg)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityChecks = accessibilityChecks
	return accessibilityInterface
}
