package accessibility

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/inclunet/accessibility/pkg/report"
)

type Inputs struct {
	Element
}

func (i *Inputs) getFieldType() string {
	value, ok := i.Selection.Attr("type")

	if ok && value != "" {
		return value
	}

	return "text"
}

func (i *Inputs) isHiddenField() bool {
	value, ok := i.Selection.Attr("type")

	if ok && value == "hidden" {
		return true
	}

	return false
}

func (i *Inputs) FindLabel() bool {
	if id, ok := i.Selection.Attr("id"); ok && id != "" {
		selection := i.Selection.ParentsUntil("form").Find("label[for=" + id + "]")

		if forId, ok := selection.Attr("for"); ok && forId == id {

			if len(selection.Text()) > 2 {
				return true
			}

		}
	}

	return false
}

func (i *Inputs) Check() (int, bool, string) {
	if i.AriaHidden() {
		return 1, true, "aria-hidden inputs do not need text description."
	}

	if i.isHiddenField() {
		return 1, true, "Hidden fields do not need description or label"
	}

	if i.FindLabel() {
		return 1, true, "This input has  a valid    label description text for screen readers."
	}

	accessibleText, ok := i.AccessibleText()

	if ok && len(accessibleText) > 3 {
		return 1, true, "This link are providing a valid    description text for screen readers."
	}

	return 1, false, "If your input field is not hidden, you need a label text description for screen reader software users."
}

func NewInputCheck(s *goquery.Selection, accessibilityReport report.AccessibilityReport) Accessibility {
	accessibilityInterface := new(Inputs)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityReport = accessibilityReport
	return accessibilityInterface
}
