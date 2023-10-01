package accessibility

type Inputs struct {
	Element
}

func (i *Inputs) AccessibleText() (string, bool) {
	switch i.getFieldType() {
	case "submit", "reset", "button":
		return i.Selection.Attr("value")
	default:
		return i.Element.GetAccessibleText()
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

func (i *Inputs) Check() ([]AccessibilityCheck, bool) {
	if i.isHiddenField() {
		return i.AddViolation("aria-hidden", false)
	}

	if i.FindLabel() {
		return i.AddViolation("emag-6.2.1", false)
	}

	accessibleText, ok := i.AccessibleText()

	if !ok {
		return i.AddViolation("emag-6.2.1", false)
	}

	return NewAccessibleTextCheck(i.AccessibilityChecks, i.AccessibilityCheck).SetMaxLength(120, "too-long-text").Check(accessibleText)
}
