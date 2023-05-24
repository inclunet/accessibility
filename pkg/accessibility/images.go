package accessibility

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type Images struct {
	Element
}

func (i *Images) isValidAlternativeDescription() bool {
	if accessibleText, ok := i.AccessibleText(); ok && len(accessibleText) >= 3 {
		return true
	}
	return false
}

func (i *Images) Check() (int, bool, string) {
	description := i.isValidAlternativeDescription()
	hidden := i.AriaHidden()

	if !description && !hidden {
		return 1, false, "No hidden imagens needs a descriptionfor accessibility"
	}

	return 1, true, "There is no errors on your image alternative text description."
}

func NewImageCheck(s *goquery.Selection, accessibilityReport report.AccessibilityReport) Accessibility {
	accessibilityInterface := new(Images)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityReport = accessibilityReport
	return accessibilityInterface
}
