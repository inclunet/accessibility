package accessibility

import (
	"github.com/PuerkitoBio/goquery"
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

func (i *Inputs) Check() AccessibilityCheck {
	accessibilityCheck := i.NewAccessibilityCheck(1, "If your input field is not hidden, you need a label text description for screen reader software users.")

	if i.AriaHidden() {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "aria-hidden inputs do not need text description."
		return accessibilityCheck
	}

	if i.isHiddenField() {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Hidden fields do not need description or label"
		return accessibilityCheck
	}

	if i.FindLabel() {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "This input has  a valid    label description text for screen readers."
		return accessibilityCheck
	}

	accessibleText, ok := i.AccessibleText()

	if ok {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Short descriptions are not a good accessiblity practice for images"

		if len(accessibleText) > 3 {
			accessibilityCheck.Description = "This link are providing a valid    description text for screen readers."
		}
	}

	return accessibilityCheck
}

func NewInputCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) Accessibility {
	accessibilityInterface := new(Inputs)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityChecks = accessibilityChecks
	return accessibilityInterface
}
