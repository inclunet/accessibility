package accessibility

import (
	"github.com/PuerkitoBio/goquery"
)

type Buttons struct {
	Element
}

func (b *Buttons) Check() AccessibilityCheck {
	accessibilityCheck := b.NewAccessibilityCheck(1, "Hidden buttons do not need accessibility description.")
	accessibilityCheck.Pass = true

	if !b.AriaHidden() {
		accessibilityCheck.Pass = false
		accessibilityCheck.Description = "No hidden buttons needs a accessible description text to screen readers"
		accessibleText, ok := b.AccessibleText()

		if !ok {
			accessibilityCheck, err := b.DeepCheck(b.Selection.Children(), b.AccessibilityChecks)
			if err == nil {
				return accessibilityCheck
			}
		}

		if ok {
			accessibilityCheck.Pass = true
			accessibilityCheck.Description = "Short descriptions do not provid good information to blind people."

			if len(accessibleText) > 3 {
				accessibilityCheck.Description = "This button are providing a valid    description text for screen readers."
			}
		}
	}

	return accessibilityCheck
}

func NewButtonCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) Accessibility {
	accessibilityInterface := new(Buttons)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityChecks = accessibilityChecks
	return accessibilityInterface
}
