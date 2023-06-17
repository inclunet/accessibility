package accessibility

type Links struct {
	Element
}

func (l *Links) Check() AccessibilityCheck {
	accessibilityCheck := l.NewAccessibilityCheck(1, "no hidden links needs a accessible text description for screen reader software.")

	if l.AriaHidden() {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Hidden Links do not need an accessible text description for screen readers"
		return accessibilityCheck
	}

	accessibleText, ok := l.AccessibleText()

	if !ok {
		accessibilityCheck, err := l.DeepCheck(l.Selection.Children(), l.AccessibilityChecks)

		if err == nil {
			return accessibilityCheck
		}
	}

	if ok {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Short alternative text       do not provide good for screen readers."

		if len(accessibleText) > 3 {
			accessibilityCheck.Description = "This link are providing a valid    description text for screen readers."
		}
	}

	return accessibilityCheck
}
