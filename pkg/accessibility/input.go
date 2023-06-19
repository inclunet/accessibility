package accessibility

type Inputs struct {
	Element
}

func (i *Inputs) AccessibleText() (string, bool) {
	switch i.getFieldType() {
	case "submit", "reset", "button":
		return i.Selection.Attr("value")
	default:
		return i.Element.AccessibleText()
	}
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
	accessibilityCheck := i.NewAccessibilityCheck("pass")

	if i.AriaHidden() {
		return i.FindViolation(accessibilityCheck, "aria-hidden")
	}

	if i.isHiddenField() {
		return i.FindViolation(accessibilityCheck, "aria-hidden")
	}

	if i.FindLabel() {
		return i.FindViolation(accessibilityCheck, "emag-6.2.1")
	}

	accessibleText, ok := i.AccessibleText()

	if !ok {
		return i.FindViolation(accessibilityCheck, "emag-6.2.1")
	}

	if i.CheckTooShortText(accessibleText) {
		return i.FindViolation(accessibilityCheck, "too-short-text")
	}

	if i.CheckTooLongText(accessibleText, 120) {
		return i.FindViolation(accessibilityCheck, "too-long-text")
	}

	return accessibilityCheck
}
