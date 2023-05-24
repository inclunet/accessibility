package accessibility

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type Links struct {
	Element
}

func (l *Links) Check() (int, bool, string) {
	if !l.AriaHidden() {
		accessibleText, ok := l.AccessibleText()

		if !ok {
			A, Pass, Description, err := DeepCheck(l.Selection.Children(), l.AccessibilityReport)
			if err == nil {
				return A, Pass, Description
			}

		}

		if ok && len(accessibleText) > 3 {
			return 1, true, "This link are providing a valid    description text for screen readers."
		}

		return 1, false, "If your link is not hidden, you need a text description for screen reader software."
	}
	return 1, true, "Hidden Links do not need text description."
}

func NewLinkCheck(s *goquery.Selection, accessibilityReport report.AccessibilityReport) Accessibility {
	accessibilityInterface := new(Links)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityReport = accessibilityReport
	return accessibilityInterface
}
