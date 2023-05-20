package accessibility

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type Buttons struct {
	Element
}

func (b *Buttons) Check() (int, bool, string) {
	if !b.AriaHidden() {
		accessibleText, ok := b.AccessibleText()

		if !ok {
			A, Pass, Description, err := DeepCheck(b.Selection.Children(), b.AccessibilityReport)
			if err == nil {
				return A, Pass, Description
			}

		}

		if ok && len(accessibleText) > 3 {
			return 1, true, "This button are providing a valid    description text for screen readers."
		}

		return 1, false, "If your button is not hidden, you need a text description for screen reader software."
	}
	return 1, true, "Hidden buttons do not need text description."
}

func NewButtonCheck(s *goquery.Selection, accessibilityReport *report.AccessibilityReport) Accessibility {
	accessibilityInterface := new(Buttons)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityReport = accessibilityReport
	return accessibilityInterface
}
