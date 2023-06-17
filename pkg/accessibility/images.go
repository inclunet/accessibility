package accessibility

type Images struct {
	Element
}

func (i *Images) Check() AccessibilityCheck {
	accessibilityCheck := i.NewAccessibilityCheck(1, "No aria-hidden images needs a valid accessibility text description.")

	if i.AriaHidden() {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Aria-hidden images do not need a valid accessibility text description"
		return accessibilityCheck
	}

	if accessibleText, ok := i.AccessibleText(); ok {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Please verify if your image has a good description"

		if len(accessibleText) >= 3 {
			accessibilityCheck.Description = "This image contains a good descriptiion"
		}
	}

	return accessibilityCheck
}
