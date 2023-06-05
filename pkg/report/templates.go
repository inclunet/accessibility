package report

import "github.com/inclunet/accessibility/pkg/accessibility"

func FilterChecks(accessibilityChecks []accessibility.AccessibilityCheck, element string) []accessibility.AccessibilityCheck {
	newAccessibilityChecks := []accessibility.AccessibilityCheck{}

	for _, accessibilityCheck := range accessibilityChecks {
		if accessibilityCheck.Element == element {
			newAccessibilityChecks = append(newAccessibilityChecks, accessibilityCheck)
		}
	}

	return newAccessibilityChecks
}
